package utilities

import (
  "pigeon/infrastructure/aws"
)

var source_addr string
var region string

func SetSourceAddr(source string) {
	aws.SetSourceAddr(source)
}

func SetRegion(reg string) {
	aws.SetRegion(reg)
}
func SendEmail(to []string, html string, subj string) {
  aws.SendEmail(to, html, subj)
}
