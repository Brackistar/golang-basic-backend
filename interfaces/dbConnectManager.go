package interfaces

import "context"

type DBConnectManager interface {
	GetDbName() string
	Connect(context.Context) error
	IsConnected() bool
	GetDataOrigin() DataOrigin
}

type DataOrigin interface {
	CreateRecord(source any, args ...any) (any, bool, error)
	GetRecord(source any, args ...any) (any, error)
	UpdateRecord(source any, args ...any) error
	DeleteRecord(source any, args ...any) error
}
