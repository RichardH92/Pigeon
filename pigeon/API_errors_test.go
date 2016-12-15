package main

import (
	"testing"
)

func TestGetInvalidQueryParameterMessageCodeIsCorrect(t *testing.T) {
	e := GetInvalidQueryParameterMessage("test")

	if e.Code != 401 {
		t.Fail()
	}
}

func TestGetInvalidQueryParameterMessageIsCorrect(t *testing.T) {
	e := GetInvalidQueryParameterMessage("test")

	temp := e.Messages.Front()
	msg := temp.Value

	if msg != "Invalid value for query parameter test" {
		t.Fail()
	}
}

func TestGetInvalidQueryParameterMessageListHasLengthOne(t *testing.T) {
	e := GetInvalidQueryParameterMessage("test")

	if e.Messages.Len() != 1 {
		t.Fail()
	}
}

func TestGetInvalidQueryParametersMessageListHasCorrectLength(t *testing.T) {
	e := GetInvalidQueryParametersMessage([]string{"test1", "test2", "test3"})

	if e.Messages.Len() != 3 {
		t.Fail()
	}
}

func TestGetInvalidQueryParametersMessagesIsCorrect(t *testing.T) {
	e := GetInvalidQueryParametersMessage([]string{"test1", "test2", "test3"})

	temp1 := e.Messages.Front()
	temp2 := temp1.Next()
	temp3 := temp2.Next()

	msg1 := temp1.Value
	msg2 := temp2.Value
	msg3 := temp3.Value

	if msg1 != "Invalid value for query parameter test1" {
		t.Fail()
	}
	if msg2 != "Invalid value for query parameter test2" {
		t.Fail()
	}
	if msg3 != "Invalid value for query parameter test3" {
		t.Fail()
	}
}

func TestGetInvalidQueryParametersMessageCodeIsCorrect(t *testing.T) {
	e := GetInvalidQueryParametersMessage([]string{"test1", "test2", "test3"})

	if e.Code != 401 {
		t.Fail()
	}
}

func TestGetFileOpeningErrorMessageCodeIsCorrect(t *testing.T) {
	e := GetFileOpeningErrorMessage("test")

	if e.Code != 501 {
		t.Fail()
	}

}
func TestGetFileOpeningErrorMessageMessageIsCorrect(t *testing.T) {
	e := GetFileOpeningErrorMessage("test")

	temp := e.Messages.Front()
	msg := temp.Value

	if msg != "Error opening file test" {
		t.Fail()
	}
}
