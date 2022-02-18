package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	Sent(data []byte) error
}

type HandleFunc func(*net.TCPConn, []byte, int) error
