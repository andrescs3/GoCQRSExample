package kafka

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/segmentio/kafka-go"
)

const (
	topic         = "message-log"
	brokerAddress = "localhost:9092"
)

func Publish(ctx context.Context, message string, id int) {
	l := log.New(os.Stdout, "kafka writer: ", 0)
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{brokerAddress},
		Topic:   topic,
		Logger:  l,
	})
	err := w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(strconv.Itoa(id)),
		Value: []byte(message),
	})
	if err != nil {
		panic("could not write message " + err.Error())
	}
}
