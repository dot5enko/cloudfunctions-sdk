package storage

import (
	"reflect"
	"time"

	"github.com/dot5enko/cloudfunctions-sdk/errors"
	"github.com/dot5enko/cloudfunctions-sdk/stat"
)

type StorageHandler func(item interface{}) error

var StoreHandler StorageHandler

func _StoreObject(ctx stat.PerformCtx, object interface{}) error {

	t0 := time.Now()
	defer ctx.Metrics.Storage.Call(t0)

	value := reflect.ValueOf(object)
	if value.Type().Kind() != reflect.Ptr {
		return errors.PersistanceNotAddressable
	}

	return StoreHandler(object)
}

func _FetchObject(ctx stat.PerformCtx, object interface{}) error {
	t0 := time.Now()
	defer ctx.Metrics.Storage.Call(t0)

	value := reflect.ValueOf(object)
	if value.Type().Kind() != reflect.Ptr {
		return errors.PersistanceNotAddressable
	}

	return StoreHandler(object)
}
