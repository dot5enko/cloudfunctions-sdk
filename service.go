package sdk

type ServiceHandler func(ctx ServiceContext) error

type ServiceHandlerError uint

const (
	FatalErrorWithStop ServiceHandlerError = iota
)

func (e ServiceHandlerError) Error() string {
	return Sprintf("service error : %d", e)
}

type ServiceContext struct {
	Data map[string]interface{}
}
