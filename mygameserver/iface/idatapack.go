package iface

type IDataPack interface {
	GetHeaderLength() uint32
	Pack(msg IMessage) ([]byte, error)
	UnpackHeader([]byte) (Header, error)
}
