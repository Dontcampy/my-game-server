package net

import (
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"io"
	"net"
	"reflect"
	"sync"
	"testing"
)

func handleData(t testing.TB, wg *sync.WaitGroup, server net.Listener, expectedMessages []Message) {
	defer wg.Done()
	conn, err := server.Accept()
	if err != nil {
		panic(err)
	}
	dp := NewDataPack()

	for i := 0; i < 2; i += 1 {
		// Unpack header
		headerData := make([]byte, dp.GetHeaderLength())
		_, err := io.ReadFull(conn, headerData)
		if err != nil {
			panic(err)
		}
		header, err := dp.UnpackHeader(headerData)
		if err != nil {
			panic(err)
		}

		// Assert contentLength
		expectedContentLength := expectedMessages[header.Id].header.ContentLength
		if header.ContentLength != expectedContentLength {
			t.Errorf("Header.ContentLength got %d, but want %d", header.ContentLength, expectedContentLength)
		}

		// Assert Id
		expectedId := expectedMessages[header.Id].header.Id
		if header.Id != expectedId {
			t.Errorf("Header.Id got %d, but want %d", header.Id, expectedId)
		}

		// Assert content
		content := make([]byte, header.ContentLength)
		_, err = io.ReadFull(conn, content)
		if err != nil {
			panic(err)
		}
		expectedContent := expectedMessages[header.Id].content
		if !reflect.DeepEqual(content, expectedContent) {
			t.Errorf("got %v want %v", content, expectedContent)
		}
		fmt.Printf("Received header: %v, content: %s\n", header, content)
	}
}

func TestDataPack_Pack(t *testing.T) {
	content1 := []byte("Hello, World.")
	header1 := iface.Header{ContentLength: uint32(len(content1)), Id: 0}
	content2 := []byte("Good night.")
	header2 := iface.Header{ContentLength: uint32(len(content2)), Id: 1}
	expectedMessages := []Message{
		{header1, content1},
		{header2, content2},
	}

	server, err := net.Listen("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go handleData(t, &wg, server, expectedMessages)

	client, err := net.Dial("tcp", "127.0.0.1:7777")
	if err != nil {
		panic(err)
	}
	dp := NewDataPack()
	pack1, err := dp.Pack(&expectedMessages[0])
	if err != nil {
		panic(err)
	}
	pack2, err := dp.Pack(&expectedMessages[1])
	if err != nil {
		panic(err)
	}

	// merge two pack
	pack1 = append(pack1, pack2...)

	_, err = client.Write(pack1)
	if err != nil {
		panic(err)
	}

	wg.Wait()
}
