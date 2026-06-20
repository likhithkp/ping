package dto

import _const "github.com/likhithkp/ping/utils/const"

type Message struct {
	Id           string             `json:"id,omitempty"`
	ChannelId    string             `json:"channelId,omitempty"`
	SenderId     string             `json:"senderId,omitempty"`
	AckMessageId string             `json:"ackMessageId,omitempty"`
	Message      string             `json:"message,omitempty"`
	Type         _const.MessageType `json:"type,omitempty"`
}
