package sdk

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
}

func (ci *CloudfunctionContext) NewTask(name string, value Task) error {

	workerTube, ok := ci.Workers[name]
	if !ok {
		return NoSuchWorkerQueue
	}

	workerTube.Push(value)

	return nil

}

type ApiHandler func(ctx CloudfunctionContext) error
