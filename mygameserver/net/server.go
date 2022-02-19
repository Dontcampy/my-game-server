package net

import (
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"github.com/dontcampy/my-game-server/mygameserver/utils"
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
		Name:      utils.GlobalObject.Name,
		IPVersion: "tcp4",
		IP:        utils.GlobalObject.Host,
		Port:      utils.GlobalObject.TcpPort,
		Router:    nil,
	}
	utils.GlobalObject.TcpServer = s
	return s
}

func (s *Server) Start() {
	fmt.Printf(
		"[MyServer] Server Name: %s, listeneer at IP: %s, Port:%d\n",
		utils.GlobalObject.Name,
		utils.GlobalObject.Host,
		utils.GlobalObject.TcpPort,
	)
	fmt.Printf(
		"[MyServer] Version %s, MaxConn: %d, MaxPackageSize: %d\n",
		utils.GlobalObject.Version,
		utils.GlobalObject.MaxConn,
		utils.GlobalObject.MaxPackageSize,
	)

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
