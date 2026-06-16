package notification

import (
	"os"
	"path"

	"github.com/likhithkp/ping/utils/ctx"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/messaging"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

func NewFcmClient(logger *zap.Logger) (*messaging.Client, error) {
	wd, _ := os.Getwd()
	saPath := path.Join(wd, "/fitnear-firebase-creds.json")
	sa := option.WithCredentialsFile(saPath)
	app, err := firebase.NewApp(ctx.Background, nil, sa)
	if err != nil {
		logger.Error("Failed to create firebase app", zap.Error(err))
		return nil, err
	}
	client, err := app.Messaging(ctx.Background)
	if err != nil {
		logger.Error("Failed to create fcm client", zap.Error(err))
		return nil, err
	}
	return client, nil
}
