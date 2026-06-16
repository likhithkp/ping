package storage

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"github.com/likhithkp/ping/utils/config"
)

type Uploader struct {
	s3Client *s3.Client
	env      *config.Env
}

func NewUploader(s3Client *s3.Client, env *config.Env) *Uploader {
	return &Uploader{
		s3Client: s3Client,
		env:      env,
	}
}

func (u *Uploader) UploadFile(ctx context.Context, file io.Reader, key string, contentType string) (string, error) {
	_, err := u.s3Client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(u.env.S3BucketName),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(contentType),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.env.S3BucketName, u.env.AwsRegion, key)
	return url, nil
}

func (u *Uploader) UploadFiles(ctx context.Context, files []io.Reader, prefix string, contentType string) ([]string, error) {
	var urls []string

	for _, file := range files {
		uniqueKey := fmt.Sprintf("%s-%s-%d", prefix, uuid.NewString(), time.Now().UnixNano())
		url, err := u.UploadFile(ctx, file, uniqueKey, contentType)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	return urls, nil
}

func (u *Uploader) DeleteFile(ctx context.Context, key string) error {
	_, err := u.s3Client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(u.env.S3BucketName),
		Key:    aws.String(key),
	})
	return err
}
