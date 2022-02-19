package iface

// IServer Server interface
type IServer interface {
	Start()
	Stop()
	Serve()
	AddRouter(router IRouter)
}
