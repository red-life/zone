package mq

import (
	"context"
	"github.com/redis/go-redis/v9"
)

var _ MessageQueue = (*RedisMessageQueue)(nil)

func NewRedisMessageQueue(rdb *redis.Client) *RedisMessageQueue {
	return &RedisMessageQueue{
		rdb: rdb,
	}
}

type RedisMessageQueue struct {
	rdb *redis.Client
}

func (mq *RedisMessageQueue) Publish(ctx context.Context, channel string, msg string) error {
	err := mq.rdb.Publish(ctx, channel, msg)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (mq *RedisMessageQueue) Subscribe(ctx context.Context, channel string) func(context.Context) (Message, error) {
	sub := mq.rdb.Subscribe(ctx, channel)
	return func(ctx context.Context) (Message, error) {
		msg, err := sub.ReceiveMessage(ctx)
		if err != nil {
			return Message{}, err
		}
		return Message{
			Payload: msg.Payload,
			Channel: msg.Channel,
		}, nil
	}

}
