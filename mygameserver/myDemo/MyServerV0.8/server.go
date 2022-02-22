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
	fmt.Println("Call PingRouter Handle...")
	fmt.Println("recv from client: id = ", request.GetHeader().Id, "content = ", string(request.GetContent()))
	err := request.GetConnection().SendMessage(200, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloWorldRouter struct {
	net.BaseRouter
}

func (p *HelloWorldRouter) Handle(request iface.IRequest) {
	fmt.Println("Call HelloWorldRouter Handle...")
	fmt.Println("recv from client: id = ", request.GetHeader().Id, "content = ", string(request.GetContent()))
	err := request.GetConnection().SendMessage(201, []byte("Hello world."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := net.NewServer(utils.GlobalObject.Name)
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloWorldRouter{})
	s.Serve()
}
