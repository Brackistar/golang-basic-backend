package interfaces

import "context"

type ConfigurationManager[T any] interface {
	InitConfig()
	GetContext() *context.Context
	GetConfig() T
}
