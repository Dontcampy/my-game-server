package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	SendMessage(id uint32, content []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
