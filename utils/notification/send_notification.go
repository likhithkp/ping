package notification

import (
	"fmt"
	"time"

	_const "github.com/likhithkp/ping/utils/const"
	"github.com/likhithkp/ping/utils/ctx"

	"firebase.google.com/go/v4/messaging"
	"go.uber.org/zap"
)

type SendNotificationUtil struct {
	client *messaging.Client
	logger *zap.Logger
}

func NewSendNotificationUtil(client *messaging.Client, logger *zap.Logger) *SendNotificationUtil {
	return &SendNotificationUtil{
		client: client,
		logger: logger,
	}
}

func (h *SendNotificationUtil) SendNotificationGoroutine(deviceToken string, title string, body string, data map[string]string, deviceType _const.DeviceType, priority _const.NotificationPriorityType, notificationTag *string, ttl *time.Duration) {
	go func() {
		h.sendNotification(deviceToken, title, body, data, deviceType, priority, notificationTag, ttl)
	}()
}

func (h *SendNotificationUtil) sendNotification(
	deviceToken string,
	title string,
	body string,
	data map[string]string,
	deviceType _const.DeviceType,
	priority _const.NotificationPriorityType,
	notificationTag *string,
	ttl *time.Duration) error {
	message := &messaging.Message{
		Data: data,
		Notification: &messaging.Notification{
			Title: title,
			Body:  body,
		},
		Token: deviceToken,
	}

	switch deviceType {
	case _const.DEVICE_TYPE_ANDROID:
		androidPriority := "normal"
		if priority == _const.NOTIFICATION_PRIORITY_TYPE_HIGH {
			androidPriority = "high"
		}
		message.Android = &messaging.AndroidConfig{
			Priority: androidPriority,
			Data:     data,
			Notification: &messaging.AndroidNotification{
				Title: title,
				Body:  body,
				Sound: "default",
			},
		}
		if ttl != nil {
			message.Android.TTL = ttl
		}
		if notificationTag != nil {
			message.Android.Notification.Tag = *notificationTag
		}
	case _const.DEVICE_TYPE_IOS:
		customData := map[string]interface{}{}
		for key, value := range data {
			customData[key] = value
		}
		apnsPriority := "5"
		if priority == _const.NOTIFICATION_PRIORITY_TYPE_HIGH {
			apnsPriority = "10"
		}
		headers := map[string]string{
			"apns-priority": apnsPriority,
		}
		if ttl != nil {
			headers["apns-expiration"] = fmt.Sprintf("%d", time.Now().Add(*ttl).Unix())
		}
		message.APNS = &messaging.APNSConfig{
			Headers: headers,
			Payload: &messaging.APNSPayload{
				CustomData: customData,
				Aps: &messaging.Aps{
					CustomData: customData,
					Alert: &messaging.ApsAlert{
						Title: title,
						Body:  body,
					},
					Sound:            "default",
					ContentAvailable: true,
				},
			},
		}
		if notificationTag != nil {
			message.APNS.Payload.Aps.ThreadID = *notificationTag
		}
	}

	_, err := h.client.Send(ctx.Background, message)
	if err != nil {
		h.logger.Error("send notification failed", zap.Error(err))
	} else {
		h.logger.Info("notification sent successfully", zap.String("token", deviceToken))
	}

	return err
}
