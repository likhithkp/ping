package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/ses"
)

func NewAWSConfig() (aws.Config, error) {
	return config.LoadDefaultConfig(context.TODO())
}

func NewS3Client(cfg aws.Config) *s3.Client {
	return s3.NewFromConfig(cfg)
}

func NewSesClient(cfg aws.Config) *ses.Client {
	return ses.NewFromConfig(cfg)
}
