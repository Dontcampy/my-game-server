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

	// Waiting for client.
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err", err)
			continue
		}

		// Simple write-back.
		go func() {
			for {
				buf := make([]byte, 512)
				// Read bytes from client.
				cnt, err := conn.Read(buf)
				if err != nil {
					fmt.Println("rece buf err", err)
					continue
				}

				fmt.Printf("recv client buf %s, cnt %d\n", buf, cnt)

				// Write back bytes to client.
				if _, err := conn.Write(buf[:cnt]); err != nil {
					fmt.Println("write back buf err", err)
					continue
				}
			}
		}()
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
