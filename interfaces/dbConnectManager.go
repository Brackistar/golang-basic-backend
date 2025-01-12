package interfaces

import "context"

type DBConnectManager interface {
	GetDbName() string
	Connect(context.Context) error
	IsConnected() bool
}
