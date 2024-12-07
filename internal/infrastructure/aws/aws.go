package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/kaitokid2302/NewsAI/internal/config"
)

func AwsInit() *session.Session {
	newSession, er := session.NewSession(&aws.Config{
		Region:      &config.Global.Region,
		Credentials: credentials.NewStaticCredentials(config.Global.PublicAccessKey, config.Global.PrivateAccessKey, ""),
	})
	if er != nil {
		panic(er)
	}
	return newSession
}
