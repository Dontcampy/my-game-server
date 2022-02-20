package iface

type IMessageHandler interface {
	DispatchHandler(request IRequest)
	AddRouter(messageId uint32, router IRouter)
}
