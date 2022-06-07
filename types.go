package sdk

type Request interface {
	Get(name string) *string
}

type Client struct {
}

type CloudfunctionContext struct {
	Request  Request
	Client   *Client
	Response map[string]interface{}
}

func (ci *CloudfunctionContext) NewTask(name string, value Task) error {
	return nil
}

type ApiHandler func(ctx CloudfunctionContext) error
