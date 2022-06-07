package sdk

type Task struct {
}

type WorkerHandler func(item Task) error
