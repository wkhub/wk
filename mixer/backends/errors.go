package backends

type backendError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *backendError) Error() string {
	return error.msg
}

func BackendError(msg string) error {
	return &backendError{msg}
}
