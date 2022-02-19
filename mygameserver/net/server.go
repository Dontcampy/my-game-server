package net

import (
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
	Router    iface.IRouter
}

// NewServer /*
func NewServer(name string) iface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      8999,
		Router:    nil,
	}
	return s
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
		dealConn := NewConnection(conn, cid, s.Router)
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

func (s *Server) AddRouter(router iface.IRouter) {
	s.Router = router
	fmt.Println("Add router successfully.")
}
