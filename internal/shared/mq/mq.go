package mq

import "context"

type MessageQueue interface {
	Publish(ctx context.Context, channel string, message string) error
	Subscribe(ctx context.Context, channel string) func(context.Context) (Message, error)
}

type Message struct {
	Payload string
	Channel string
}
