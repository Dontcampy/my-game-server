package net

import (
	"errors"
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"io"
	"net"
)

type Connection struct {
	Conn           *net.TCPConn
	ConnID         uint32
	isClosed       bool
	ExitChan       chan bool
	MessageHandler iface.IMessageHandler
}

func NewConnection(conn *net.TCPConn, connID uint32, messageHandler iface.IMessageHandler) *Connection {
	c := &Connection{
		Conn:           conn,
		ConnID:         connID,
		isClosed:       false,
		ExitChan:       make(chan bool, 1),
		MessageHandler: messageHandler,
	}

	return c
}

func (c *Connection) StartReader() {
	fmt.Println("Read")
	defer fmt.Println("connID = ", c.ConnID, "Reader is exit, remote addr is ", c.RemoteAddr().String())
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
		go c.MessageHandler.DispatchHandler(&req)
	}
}

func (c *Connection) Start() {
	fmt.Println("Conn Start()... ConnID = ", c.ConnID)
	go c.StartReader()
}

func (c *Connection) Stop() {
	fmt.Println("Conn Stop().. ConnID = ", c.ConnID)

	if c.isClosed {
		return
	}
	c.isClosed = true

	c.Conn.Close()
	close(c.ExitChan)
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

	// send pack
	_, err = c.Conn.Write(pack)
	if err != nil {
		return err
	}

	return nil
}
