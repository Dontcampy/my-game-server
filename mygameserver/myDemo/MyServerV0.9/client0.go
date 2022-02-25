package main

import (
	"fmt"
	net2 "github.com/dontcampy/my-game-server/mygameserver/net"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		dp := net2.NewDataPack()
		pack, err := dp.Pack(net2.NewMessage(0, []byte("MyServerV0.8 Client Test Message")))
		if err != nil {
			panic(err)
		}
		_, err = conn.Write(pack)
		if err != nil {
			panic(err)
		}

		rawHead := make([]byte, dp.GetHeaderLength())
		_, err = io.ReadFull(conn, rawHead)
		if err != nil {
			panic(err)
		}
		header, err := dp.UnpackHeader(rawHead)
		if err != nil {
			panic(err)
		}

		var content []byte
		if header.ContentLength > 0 {
			content = make([]byte, header.ContentLength)
			_, err = io.ReadFull(conn, content)
			if err != nil {
				panic(err)
			}
		}
		fmt.Println("recv from server: id = ", header.Id, "content = ", string(content))

		// blocking cpu
		time.Sleep(1 * time.Second)
	}
}
