package service

import (
	"context"
	"encoding/json"
	"fmt"
	"learn-microservices-mail/dto"
	"time"

	"github.com/segmentio/kafka-go"
)

type kafkaService struct {
	reader *kafka.Reader
	// writer      *kafka.Writer
	mailService MailService
}

type KafkaService interface {
	ListenTopic() error
}

func NewKafkaService(ms MailService) (*kafkaService, error) {
	topic := "send-mail"
	partition := 0
	host := "localhost:9092"

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{host},
		Topic:     topic,
		GroupID:   "mail-server",
		Partition: partition,
		MaxBytes:  10e6,
		MaxWait:   1 * time.Second,
	})
	r.SetOffset(42)

	// w := &kafka.Writer{
	// 	Addr:     kafka.TCP(host),
	// 	Topic:    topic,
	// 	Balancer: &kafka.LeastBytes{},
	// }

	return &kafkaService{
		reader: r,
		// writer:      w,
		mailService: ms,
	}, nil
}

func (k *kafkaService) ListenTopics() error {
	fmt.Println("Starting listen topics...")
	defer k.reader.Close()
	for {
		m, err := k.reader.ReadMessage(context.Background())
		if err != nil {
			return err
		}

		var msgData dto.MailTransferData
		err = json.Unmarshal(m.Value, &msgData)
		if err != nil {
			return err
		}

		// go k.mailService.SendMail(msgData.To, msgData.Subject, msgData.Template)

		fmt.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))
	}
}
