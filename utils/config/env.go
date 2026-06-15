package config

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_const "github.com/likhithkp/ping/utils/const"
)

type Env struct {
	DeploymentEnv string
	MongoUri      string
	Database      string
	Port          string
	JwtSecret     string
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

	return &Env{
		MongoUri:  databaseUri,
		Database:  database,
		Port:      port,
		JwtSecret: jwtSecret,
	}, nil
}
