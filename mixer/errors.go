package mixer

// Define your Error struct
type mixerError struct {
	msg string
}

// Create a function Error() string and associate it to the struct.
func (error *mixerError) Error() string {
	return error.msg
}

func MixerError(msg string) error {
	return &mixerError{msg}
}
