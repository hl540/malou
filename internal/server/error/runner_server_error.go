package error

type Error struct {
	Code    int32
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
