package iface

type IConnectionManager interface {
	Add(connection IConnection)
	Remove(connId uint32)
	Get(connId uint32) (IConnection, error)
	Size() int
	Clear()
}
