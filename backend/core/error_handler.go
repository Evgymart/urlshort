package core

type ErrorMessage struct {
	Error string
}

func HandleError(err error) ErrorMessage {
	return ErrorMessage{err.Error()}
}
