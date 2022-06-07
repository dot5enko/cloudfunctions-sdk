package sdk

type ServiceHandler func(ctx ServiceContext) error

type ServiceContext struct {
	Data map[string]interface{}
}
