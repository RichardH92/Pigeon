package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
)

var source_addr string
var region string

func SetSourceAddr(addr string) {
	source_addr = addr
}

func SetRegion(reg string) {
	region = reg
}

func AwsSendEmail(to []string, html string, subj string) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return
	}

	dests := make([]*string, len(to))
	for i, _ := range to {
		dests[i] = &to[i]
	}

	svc := ses.New(sess, &aws.Config{Region: aws.String(region)})

	params := &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: dests,
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data:    aws.String(html),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(subj),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(source_addr),
	}
	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
