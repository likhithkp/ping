package dto

import _const "github.com/likhithkp/ping/utils/const"

type Message struct {
	Id           string `json:"id"`
	ChannelId    string `json:"channelId"`
	SenderId     string `json:"senderId"`
	AckMessageId string `json:"ackMessageId"`
	Message      string `json:"message"`
	// Status       _const.MessageStatusType `json:"status"`
	Type _const.MessageType `json:"type"`
}
