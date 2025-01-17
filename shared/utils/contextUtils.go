package utils

import (
	"context"

	"github.com/Brackistar/golang-basic-backend/shared/models"
)

func GetContextValue[T interface{}](ctx *context.Context, key models.Key) T {
	return (*ctx).Value(key).(T)
}
