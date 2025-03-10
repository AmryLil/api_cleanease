package utils

import (
	"api_cleanease/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func NewSession() (*session.Session, error) {
	cfg := config.LoadAwsConfig()
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
		Credentials: credentials.NewStaticCredentials(
			cfg.AccessKeyID,
			cfg.AccessKeySecret,
			"",
		),
	})

	if err != nil {
		return nil, err
	}

	return sess, nil
}
