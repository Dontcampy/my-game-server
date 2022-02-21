package utils

import (
	"encoding/json"
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"io/ioutil"
	"runtime"
)

type globalObj struct {
	TcpServer iface.IServer
	Host      string
	TcpPort   int
	Name      string

	Version          string
	MaxConn          int
	MaxPackageSize   uint32
	MaxTaskQueueSize uint32
	WorkerPoolSize   uint32
}

var GlobalObject *globalObj

func (g *globalObj) Reload() {
	data, err := ioutil.ReadFile("conf/config.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(data, g)
	if err != nil {
		panic(err)
	}

}

func init() {
	GlobalObject = &globalObj{
		Name:             "SiriusServerApp",
		Version:          "V0.4",
		TcpPort:          8999,
		Host:             "0.0.0.0",
		MaxConn:          1000,
		MaxPackageSize:   4096,
		MaxTaskQueueSize: 1024,
		WorkerPoolSize:   uint32(runtime.NumCPU()),
	}

	GlobalObject.Reload()
}
