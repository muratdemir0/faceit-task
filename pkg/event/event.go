package event

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
}

func NewProducer(url string) (Producer, error) {
	w := kafka.Writer{
		Addr:     kafka.TCP(url),
		Balancer: &kafka.LeastBytes{},
	}

	return Producer{writer: &w}, nil
}

func (p Producer) Produce(ctx context.Context, topic string, message interface{}) error {
	messageBytes, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return p.writer.WriteMessages(ctx, kafka.Message{
		Topic: topic,
		Value: messageBytes,
	})
}
