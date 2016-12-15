package main

import (
	//"container/list"
	"net/mail"
	"strings"
)

var email_queries []Email_Type

func SetEmailQueries(queries []Email_Type) {
	email_queries = queries
}

func ValidateDestQuery(dest string) ([]string, bool) {
	if dest == "" {
		return nil, false
	}

	dests := strings.Split(dest, "+")
	for _, element := range dests {
		if !ValidateEmailQuery(element) {
			return dests, false
		}
	}

	return dests, true
}

func ValidateEmailQuery(email string) bool {
	_, err := mail.ParseAddress(email)

	if err == nil {
		return true
	}

	return false
}

func ValidateEmailType(emailType string) (Email_Type, bool) {
	for _, element := range email_queries {

		if emailType == element.Type_Name {
			return element, true
		}
	}

	return Email_Type{}, false
}

func ValidateHTMLReplacementValues(m map[string]string) (bool, []string) {
	s := make([]string, 0, len(m))
	err_count := 0

	for key, _ := range m {
		if m[key] == "" {
			s = s[0 : err_count+1]
			s[err_count] = key
			err_count = err_count + 1
		}
	}

	return (err_count == 0), s
}
