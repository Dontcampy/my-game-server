package net

import "github.com/dontcampy/my-game-server/mygameserver/iface"

type Message struct {
	header  iface.Header
	content []byte
}

func NewMessage(id uint32, content []byte) *Message {
	return &Message{
		header:  iface.Header{ContentLength: uint32(len(content)), Id: id},
		content: content,
	}
}

func (m *Message) GetHeader() iface.Header {
	return m.header
}

func (m *Message) GetContent() []byte {
	return m.content
}
