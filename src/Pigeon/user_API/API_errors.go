package user_API

import (
	"container/list"
)

type ErrorMessage struct {
	Code int
	Messages *list.List
}

func GetInvalidQueryParameterMessage(element string) ErrorMessage {
	list := list.New()
	msg := "Invalid value for query parameter " + element
	list.PushBack(msg)
	return ErrorMessage{Code: 401, Messages: list}
}

func GetInvalidQueryParametersMessage(elements []string) ErrorMessage {
	list := list.New()
	msg := ""

	for _, element := range elements {
		msg = "Invalid value for query parameter " + element
		list.PushBack(msg)
	}

	return ErrorMessage{Code: 401, Messages: list}
}

func GetFileOpeningErrorMessage(name string) ErrorMessage {
	list := list.New()
	msg := "Error opening file " + name
	list.PushBack(msg)
	return ErrorMessage{Code: 501, Messages: list}
}
