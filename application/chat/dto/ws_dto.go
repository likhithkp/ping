package dto

import _const "github.com/likhithkp/ping/utils/const"

type Message struct {
	Id        string             `json:"id"`
	ChannelId string             `json:"channelId"`
	SenderId  string             `json:"senderId"`
	Message   string             `json:"message"`
	Type      _const.MessageType `json:"type"`
}
