package sdk

type Error struct {
	Code  uint
	Label string

	Cause *error
}

func (err *Error) Error() string {
	return Sprintf("%d : %s", err.Code, err.Label)
}
