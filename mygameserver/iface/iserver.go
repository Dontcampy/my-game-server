package iface

// IServer Server interface
type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(messageId uint32, router IRouter)
	GetConnectionManager() IConnectionManager
	SetOnConnectionStart(func(connection IConnection))
	SetOnConnectionStop(func(connection IConnection))
	CallOnConnectionStart(connection IConnection)
	CallOnConnectionStop(connection IConnection)
}
