package sdk

import (
	"github.com/dot5enko/cloudfunctions-sdk/errors"
	"github.com/dot5enko/cloudfunctions-sdk/storage"
)

type Request interface {
	Get(name string) string
}

type Client struct {
}

type WorkerQueueInterface interface {
	Chan() <-chan interface{}
	Push(item interface{})
}

type InterchangeMessage struct {
	ReqUid string `json:"req_id"`

	Uri   string `json:"uri"`
	Query string `json:"query"`
}

type InterchangeMessageResponse struct {
	ReqUid     string `json:"req_id"`
	HandleTime uint64 `json:"handle_time"`

	Data string `json:"data"`
}

type CloudfunctionContext struct {
	Request Request

	Client *Client

	Response interface{}
	Workers  map[string]WorkerQueueInterface

	Storage storage.StorageHandlers
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
