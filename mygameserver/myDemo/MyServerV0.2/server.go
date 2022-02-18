package main

import "github.com/dontcampy/my-game-server/mygameserver/net"

func main() {
	s := net.NewServer("[MyServer V0.2]")
	s.Serve()
}
