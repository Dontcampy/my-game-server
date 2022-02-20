package main

import (
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"github.com/dontcampy/my-game-server/mygameserver/net"
	"github.com/dontcampy/my-game-server/mygameserver/utils"
)

type PingRouter struct {
	net.BaseRouter
}

func (p *PingRouter) Handle(request iface.IRequest) {
	fmt.Println("Call Router Handle...")
	fmt.Println("recv from client: id = ", request.GetHeader().Id, "content = ", string(request.GetContent()))
	err := request.GetConnection().SendMessage(1, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := net.NewServer(utils.GlobalObject.Name)
	s.AddRouter(&PingRouter{})
	s.Serve()
}
