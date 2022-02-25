package net

import (
	"errors"
	"fmt"
	"github.com/dontcampy/my-game-server/mygameserver/iface"
	"sync"
)

type ConnectionManager struct {
	connections map[uint32]iface.IConnection
	lock        sync.RWMutex
}

func NewConnectionManager() *ConnectionManager {
	return &ConnectionManager{connections: make(map[uint32]iface.IConnection)}
}

func (c *ConnectionManager) Add(connection iface.IConnection) {
	c.lock.Lock()
	c.connections[connection.GetConnId()] = connection
	c.lock.Unlock()
	fmt.Println("connId = ", connection.GetConnId(),
		" add to ConnectionManager successfully, manager size = ", c.Size())
}

func (c *ConnectionManager) Remove(connId uint32) {
	c.lock.Lock()
	delete(c.connections, connId)
	c.lock.Unlock()
	fmt.Println("connId = ", connId,
		" remove from ConnectionManager successfully, manager size = ", c.Size())
}

func (c *ConnectionManager) Get(connId uint32) (iface.IConnection, error) {
	c.lock.RLock()
	connection, ok := c.connections[connId]
	c.lock.RUnlock()
	if !ok {
		return nil, errors.New("connection not found")
	}
	return connection, nil
}

func (c *ConnectionManager) Size() int {
	return len(c.connections)
}

func (c *ConnectionManager) Clear() {
	c.lock.Lock()
	for _, connection := range c.connections {
		connection.Stop()
	}
	c.connections = nil
	c.lock.Unlock()

	fmt.Println("Clear All connections successfully.")
}
