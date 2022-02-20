package net

import "github.com/dontcampy/my-game-server/mygameserver/iface"

type Request struct {
	// established connection
	conn iface.IConnection

	// data received message
	message iface.IMessage
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetContent() []byte {
	return r.message.GetContent()
}

func (r *Request) GetHeader() iface.Header {
	return r.message.GetHeader()
}
