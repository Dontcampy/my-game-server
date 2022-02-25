package iface

import "net"

type IConnection interface {
	Start()
	Stop()
	GetTCPConnection() *net.TCPConn
	GetConnId() uint32
	RemoteAddr() net.Addr
	SendMessage(id uint32, content []byte) error
	SetProperty(key string, value interface{})
	GetProperty(key string) (interface{}, error)
	RemoveProperty(key string)
}

type HandleFunc func(*net.TCPConn, []byte, int) error
