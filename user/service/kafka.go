package service

import (
	"context"
	"encoding/json"

	"github.com/segmentio/kafka-go"
)

type kafkaService struct {
	writer *kafka.Writer
}

type KafkaService interface {
	SendMessage(payload map[string]interface{}) error
}

func NewKafkaService(addr string, topic string) KafkaService {
	w := &kafka.Writer{
		Addr:     kafka.TCP(addr),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return &kafkaService{
		writer: w,
	}
}

func (k *kafkaService) SendMessage(payload map[string]interface{}) error {
	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}
	message := kafka.Message{
		Value: data,
	}
	err = k.writer.WriteMessages(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}
