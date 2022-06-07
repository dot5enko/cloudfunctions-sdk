package sdk

import (
	"github.com/dot5enko/cloudfunctions-sdk/errors"
	"github.com/dot5enko/cloudfunctions-sdk/storage"
)

type Request interface {
	Get(name string) *string
}

type Client struct {
}

type WorkerQueueInterface interface {
	Chan() <-chan interface{}
	Push(item interface{})
}

type CloudfunctionContext struct {
	Request Request

	Client *Client

	Response map[string]interface{}
	Workers  map[string]WorkerQueueInterface

	QueryHandler storage.QueryBuilderHandler
}

func (ci *CloudfunctionContext) NewTask(name string, value Task) error {

	workerTube, ok := ci.Workers[name]
	if !ok {
		return errors.NoSuchWorkerQueue
	}

	workerTube.Push(value)

	return nil

}

type ApiHandler func(ctx CloudfunctionContext) error
