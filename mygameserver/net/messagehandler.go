package net

import (
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"github.com/dontcampy/my-game-server/mygameserver/utils"
	"strconv"
)

type MessageHandler struct {
	apis         map[uint32]iface.IRouter
	taskQueues   []chan iface.IRequest
	workPoolSize uint32
}

func NewMessageHandler() *MessageHandler {
	return &MessageHandler{
		apis:         make(map[uint32]iface.IRouter),
		taskQueues:   make([]chan iface.IRequest, utils.GlobalObject.WorkerPoolSize),
		workPoolSize: utils.GlobalObject.WorkerPoolSize,
	}
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

func (m *MessageHandler) StartWorkerPool() {
	for i := 0; i < int(m.workPoolSize); i += 1 {
		m.taskQueues[i] = make(chan iface.IRequest, utils.GlobalObject.MaxTaskQueueSize)
		go m.startWorker(i, m.taskQueues[i])
	}
}

func (m *MessageHandler) startWorker(workerId int, taskQueue chan iface.IRequest) {
	fmt.Println("Worker ID = ", workerId, "is started...")
	for true {
		select {
		case request := <-taskQueue:
			m.DispatchHandler(request)
		}
	}
}

func (m *MessageHandler) SendMessageToTaskQueue(request iface.IRequest) {
	workerId := request.GetConnection().GetConnId() % m.workPoolSize
	fmt.Println("Add ConnId = ", request.GetConnection().GetConnId(), "request = ", request.GetHeader().Id, "to WorkerId = ", workerId)
	m.taskQueues[workerId] <- request
}
