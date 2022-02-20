package iface

// IServer Server interface
type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(messageId uint32, router IRouter)
}
