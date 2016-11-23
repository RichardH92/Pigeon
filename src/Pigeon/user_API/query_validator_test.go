package user_API

import (
	"testing"
)

func TestValidateDestQuerySucceeds(t *testing.T) {
	_, success := ValidateDestQuery("test@test.com")
	if !success {
		t.Fail()
	}
}

func TestValidateDestQueryFailsInvalidEmail(t *testing.T) {
	_, success := ValidateDestQuery("test")
	if success {
		t.Fail()
	}
}

func TestValidateDestQuerySucceedsiMultipleEmails(t *testing.T) {
	dests, success := ValidateDestQuery("test@test.com+test2@test.com")
	if !success && dests[0] != "test@test.com" && dests[1] != "test2@test.com" {
		t.Fail()
	}
}

func TestValidateDestQueryFailsOneValidOneInvalidEmail(t *testing.T) {
	_, success := ValidateDestQuery("test+test@test.com")
	if success {
		t.Fail()
	}
}

func TestValidateEmailTypeSucceeds(t *testing.T) {
	eType := Email_Type{
		Type_Name:      "Test",
		HTML_File_Name: "Test.html",
		Query_Vals:     []string{"test1", "test2"},
		Subject:        "TestSubj",
	}
	emailQueries := []Email_Type{eType}
	SetEmailQueries(emailQueries)

	_, success := ValidateEmailType("Test")
	if !success {
		t.Fail()
	}
}
