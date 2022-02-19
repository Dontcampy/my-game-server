package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	// connect to server
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!")
		return
	}

	for {
		// send msg to server
		_, err := conn.Write([]byte("Hello MyServer V0.2..."))
		if err != nil {
			fmt.Println("write conn err", err)
			return
		}

		// receive msg from server
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}

		fmt.Printf("server call back: %s, cnt = %d\n", buf, cnt)

		// blocking cpu
		time.Sleep(1 * time.Second)
	}
}
