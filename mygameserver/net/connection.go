package net

import (
	"errors"
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"github.com/dontcampy/my-game-server/mygameserver/utils"
	"io"
	"net"
)

type Connection struct {
	TcpServer      iface.IServer
	Conn           *net.TCPConn
	ConnID         uint32
	isClosed       bool
	ExitChan       chan bool
	MessageHandler iface.IMessageHandler
	MessageChannel chan []byte
}

func NewConnection(server iface.IServer, conn *net.TCPConn, connID uint32, messageHandler iface.IMessageHandler) *Connection {
	c := &Connection{
		TcpServer:      server,
		Conn:           conn,
		ConnID:         connID,
		isClosed:       false,
		ExitChan:       make(chan bool, 1),
		MessageHandler: messageHandler,
		MessageChannel: make(chan []byte),
	}
	server.GetConnectionManager().Add(c)
	return c
}

func (c *Connection) StartReader() {
	fmt.Println("[Reader Goroutine is running]")
	defer fmt.Println("[Reader is exit], connID = ", c.ConnID, ", remote addr is ", c.RemoteAddr().String())
	defer c.Stop()

	dp := NewDataPack()

	for {
		// Read client data to the buffer.
		//buf := make([]byte, utils.GlobalObject.MaxPackageSize)
		//_, err := c.Conn.Read(buf)
		//if err != nil {
		//	fmt.Println("recv buf err", err)
		//	continue
		//}

		// unpack header
		rawHeader := make([]byte, dp.GetHeaderLength())
		_, err := io.ReadFull(c.Conn, rawHeader)
		if err != nil {
			fmt.Println("read header err: ", err)
			break
		}
		header, err := dp.UnpackHeader(rawHeader)
		if err != nil {
			fmt.Println("unpack header err: ", err)
			break
		}

		// unpack content
		var content []byte
		if header.ContentLength > 0 {
			content = make([]byte, header.ContentLength)
			_, err = io.ReadFull(c.Conn, content)
			if err != nil {
				fmt.Println("read content err: ", err)
				break
			}
		}

		// pack recv data in to Request
		req := Request{
			conn:    c,
			message: &Message{header: header, content: content},
		}

		// call router handler
		if utils.GlobalObject.WorkerPoolSize > 0 {
			c.MessageHandler.SendMessageToTaskQueue(&req)
		} else {
			go c.MessageHandler.DispatchHandler(&req)
		}
	}
}

func (c *Connection) StartWriter() {
	fmt.Println("[Writer Goroutine is running]")
	defer fmt.Println("[Writer is exit], connID = ", c.ConnID, ", remote addr is ", c.RemoteAddr().String())

	for true {
		select {
		case data := <-c.MessageChannel:
			_, err := c.Conn.Write(data)
			if err != nil {
				fmt.Println("Send data error, ", err)
				return
			}
		case <-c.ExitChan:
			return
		}
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	go c.StartReader()
	go c.StartWriter()
	c.TcpServer.CallOnConnectionStart(c)
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	c.TcpServer.CallOnConnectionStop(c)

	c.Conn.Close()

	c.ExitChan <- true

	c.TcpServer.GetConnectionManager().Remove(c.ConnID)

	close(c.ExitChan)
	close(c.MessageChannel)
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnId() uint32 {
	return c.ConnID
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

func (c *Connection) SendMessage(id uint32, content []byte) error {
	if c.isClosed {
		return errors.New("connection closed when send message")
	}

	// pack message
	dp := NewDataPack()
	pack, err := dp.Pack(NewMessage(id, content))
	if err != nil {
		return err
	}

	c.MessageChannel <- pack
	return nil
}
