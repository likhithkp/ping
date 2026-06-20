package _const

type MessageType string

const (
	ACK     MessageType = "ack"
	MESSAGE MessageType = "message"
	PING    MessageType = "ping"
	PONG    MessageType = "pong"
)
