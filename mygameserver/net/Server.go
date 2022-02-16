package net

import "github.com/dontcampy/my-game-server/mygameserver/iface"

// Server IServer implement
type Server struct {
	Name      string
	IPVersion string
	IP        string
	Port      int
}

func (s *Server) Start() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Stop() {
	//TODO implement me
	panic("implement me")
}

func (s *Server) Serve() {
	//TODO implement me
	panic("implement me")
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
