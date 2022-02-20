package iface

type Header struct {
	ContentLength uint32
	Id            uint32
}

type IMessage interface {
	GetHeader() Header
	GetContent() []byte
}
