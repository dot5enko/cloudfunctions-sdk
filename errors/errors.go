package errors

import "github.com/dot5enko/cloudfunctions-sdk/std"

type SdkError uint

const (
	NoSuchWorkerQueue SdkError = iota
	FatalErrorWithStop
	PersistanceNotAddressable
)

func (e SdkError) Error() string {
	return std.Sprintf("sdk error : %d", e)
}

// type Error struct {
// 	Code  uint
// 	Label string

// 	Cause *error
// }

// func (err *Error) Error() string {
// 	return Sprintf("%d : %s", err.Code, err.Label)
// }
