package di

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func AWSSessionProvider() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))
}

func AWSSNSProvider(awSession *session.Session) *sns.SNS {
	return sns.New(awSession)
}
