package utils

import (
	"context"
	"testing"

	"github.com/Brackistar/golang-basic-backend/shared/models"
	"github.com/stretchr/testify/assert"
)

func Test_GetContextValue_KeyExists_StringKey(t *testing.T) {
	var ctx context.Context = context.TODO()
	var key models.Key = models.Key("testKey")

	var expectedResult string = "testString"

	ctx = context.WithValue(ctx, key, expectedResult)

	result := GetContextValue[string](&ctx, key)

	assert.Equal(t, expectedResult, result)
}

func Test_GetContextValue_MultipleValues_StringKey(t *testing.T) {
	var ctx context.Context = context.TODO()
	var testKey models.Key = models.Key("testKey1")
	var testKeyVal string = "testValue"
	var keys map[models.Key]string = map[models.Key]string{
		testKey:                testKeyVal,
		models.Key("testKey2"): "testVal2",
		models.Key("testKey3"): "testVal3",
	}

	var expectedResult string = testKeyVal

	for key, val := range keys {
		ctx = context.WithValue(ctx, key, val)
	}

	result := GetContextValue[string](&ctx, testKey)

	assert.Equal(t, expectedResult, result)
}

func Test_GetContextValue_KeyDoesntExists_StringKey(t *testing.T) {
	var ctx context.Context = context.TODO()
	var key models.Key = models.Key("testKey")
	var keyValue models.Key = models.Key("not the test string")

	ctx = context.WithValue(ctx, key, keyValue)

	assert.Panics(t, func() { _ = GetContextValue[string](&ctx, key) })
}
