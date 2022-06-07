package storage

import (
	"reflect"
	"time"

	sdk "github.com/dot5enko/cloudfunctions-sdk"
)

type StorageHandler func(item interface{}) error

var Handler StorageHandler

func _StoreObject(ctx sdk.PerformCtx, object interface{}) error {

	t0 := time.Now()
	defer ctx.Metrics.Storage.Call(t0)

	value := reflect.ValueOf(object)
	if value.Type().Kind() != reflect.Ptr {
		return sdk.PersistanceNotAddressable
	}

	return Handler(object)
}

func _FetchObject(ctx sdk.PerformCtx, object interface{}) error {
	t0 := time.Now()
	defer ctx.Metrics.Storage.Call(t0)

	value := reflect.ValueOf(object)
	if value.Type().Kind() != reflect.Ptr {
		return sdk.PersistanceNotAddressable
	}

	return Handler(object)
}
