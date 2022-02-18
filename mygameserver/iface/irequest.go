package iface

type IRequest interface {
	// GetConnection Get current connection.
	GetConnection() IConnection

	// GetData Get request data.
	GetData() []byte
}
