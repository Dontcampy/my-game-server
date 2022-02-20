package net

import (
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"strconv"
)

type MessageHandler struct {
	apis map[uint32]iface.IRouter
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{apis: make(map[uint32]iface.IRouter)}
}

func (m *MessageHandler) DispatchHandler(request iface.IRequest) {
	id := request.GetHeader().Id
	router, ok := m.apis[id]
	if !ok {
		fmt.Println("Unknown message id: ", id)
	}
	router.PreHandle(request)
	router.Handle(request)
	router.PostHandle(request)
}

func (m *MessageHandler) AddRouter(messageId uint32, router iface.IRouter) {
	if _, ok := m.apis[messageId]; ok {
		panic("repeat app, messageId = " + strconv.Itoa(int(messageId)))
	}
	m.apis[messageId] = router
	fmt.Println("Add api messageId = ", messageId)
}
