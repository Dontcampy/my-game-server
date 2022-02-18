package net

import (
	"errors"
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"net"
)

// Server IServer implement
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

// CallBackToClient simple handleAPI
func CallBackToClient(conn *net.TCPConn, data []byte, cnt int) error {
	fmt.Println("[Conn Handler] CallBackToClient ...")
	if _, err := conn.Write(data[:cnt]); err != nil {
		fmt.Println("write back buf err", err)
		return errors.New("CallBackToClient error")
	}

	return nil
}

func (s *Server) Start() {
	fmt.Printf("[Start] Server Listening at IP: %s, Port %d, is starting\n", s.IP, s.Port)

	go s.listen()
}

// Start listen from client
func (s *Server) listen() {
	// Resolve Address.
	addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
	if err != nil {
		fmt.Println("resolve tcp addr error: ", err)
		return
	}

	// Listening resolved address.
	listener, err := net.ListenTCP(s.IPVersion, addr)
	if err != nil {
		fmt.Println("listen", s.IPVersion, "err", err)
		return
	}

	fmt.Println("Start server successfully, ", s.Name, ", listening...")
	var cid uint32 = 0

	// Waiting for client.
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		// Init connection.
		dealConn := NewConnection(conn, cid, CallBackToClient)
		cid += 1
		// Start.
		go dealConn.Start()
	}
}

func (s *Server) Stop() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Serve() {
	// start server
	s.Start()

	// blocking main thread
	select {}
}

/*
Initialize Server
*/

func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
	}
	return s
}
