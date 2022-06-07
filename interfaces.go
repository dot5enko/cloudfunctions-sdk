package sdk

type SprintfFunctionType func(message string, a ...interface{}) string
type LogFunctionType func(message string, a ...interface{})

// should be inited externally
var Sprintf SprintfFunctionType
var Log LogFunctionType
