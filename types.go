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

func (ci *CloudfunctionContext) PutIntoChannel(name string, value interface{}) {

}

type CloudfunctionHandler func(ctx CloudfunctionContext) error
