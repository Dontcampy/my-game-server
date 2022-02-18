package net

import "github.com/dontcampy/my-game-server/mygameserver/iface"

type Request struct {
	// established connection
	conn iface.IConnection

	// data received data
	data []byte
}

func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.data
}
