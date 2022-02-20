package iface

type IRequest interface {
	// GetConnection Get current connection.
	GetConnection() IConnection

	// GetContent Get request content.
	GetContent() []byte

	// GetHeader Get request header.
	GetHeader() Header
}
