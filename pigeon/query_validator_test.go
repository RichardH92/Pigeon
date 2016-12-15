package main

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

func TestValidateEmailTypeReturnsCorrectly(t *testing.T) {
	eType := Email_Type{
		Type_Name:      "Test",
		HTML_File_Name: "Test.html",
		Query_Vals:     []string{"test1", "test2"},
		Subject:        "TestSubj",
	}
	emailQueries := []Email_Type{eType}
	SetEmailQueries(emailQueries)

	ret, _ := ValidateEmailType("Test")
	if ret.HTML_File_Name != "Test.html" {
		t.Fail()
	}
	if ret.Subject != "TestSubj" {
		t.Fail()
	}
	if len(ret.Query_Vals) != 2 {
		t.Fail()
	}
	if ret.Query_Vals[0] != "test1" {
		t.Fail()
	}
	if ret.Query_Vals[1] != "test2" {
		t.Fail()
	}
}
