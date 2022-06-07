package sdk

import "time"

type ServiceDefintion struct {
	Name     string
	Handler  ServiceHandler
	Interval time.Duration
}

type WorkerDefinition struct {
	Handler WorkerHandler
	Name    string
	Size    uint
	Threads uint
}

type AppDefinition struct {
	Services   []ServiceDefintion
	Workers    map[string]WorkerDefinition
	Handlers   map[string]ApiHandler
	Persistant []interface{}
}
