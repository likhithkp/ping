package config

import (
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_const "github.com/likhithkp/ping/utils/const"
)

type Env struct {
	DeploymentEnv      string
	MongoUri           string
	Database           string
	Port               string
	JwtSecret          string
	S3BucketName       string
	SenderEmail        string
	AwsRegion          string
	AwsAccessKey       string
	AwsSecretAccessKey string
}

func NewEnv() (*Env, error) {
	deploymentEnv := strings.TrimSpace(os.Getenv("DEPLOYMENT_ENV"))
	if deploymentEnv != string(_const.Deployment_Production) {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	databaseUri := os.Getenv("DATABASE_URI")
	if databaseUri == "" {
		log.Fatalln("[env.go] DATABASE_URI missing")
	}

	database := os.Getenv("DATABASE")
	if database == "" {
		log.Fatalln("[env.go] DATABASE missing")
	}

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatalln("[env.go] PORT missing")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatalln("[env.go] JWT_SECRET missing")
	}

	s3BucketName := os.Getenv("S3_BUCKET")
	if jwtSecret == "" {
		log.Fatalln("[env.go] S3_BUCKET missing")
	}

	senderEmail := os.Getenv("SENDER_EMAIL")
	if senderEmail == "" {
		log.Fatalln("[env.go] SENDER_EMAIL missing")
	}

	awsRegion := os.Getenv("AWS_REGION")
	if len(awsRegion) == 0 {
		return nil, errors.New("AWS_REGION is empty")
	}
	awsAccessKey := os.Getenv("AWS_ACCESS_KEY")
	if len(awsAccessKey) == 0 {
		return nil, errors.New("AWS_ACCESS_KEY is empty")
	}
	awsSecretAccessKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	if len(awsSecretAccessKey) == 0 {
		return nil, errors.New("AWS_SECRET_ACCESS_KEY is empty")
	}

	return &Env{
		MongoUri:           databaseUri,
		Database:           database,
		Port:               port,
		JwtSecret:          jwtSecret,
		S3BucketName:       s3BucketName,
		SenderEmail:        senderEmail,
		AwsRegion:          awsRegion,
		AwsAccessKey:       awsAccessKey,
		AwsSecretAccessKey: awsSecretAccessKey,
	}, nil
}
