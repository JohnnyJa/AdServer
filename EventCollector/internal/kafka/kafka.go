package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/JohnnyJa/AdServer/EventCollector/internal/model"
	"github.com/JohnnyJa/AdServer/EventCollector/service"
	"github.com/sirupsen/logrus"
)

type Kafka interface {
	service.Service
	Write(event model.Event) error
}

type kafka struct {
	Config   *Config
	Logger   *logrus.Logger
	producer sarama.AsyncProducer
}

func New(config *Config, logger *logrus.Logger) Kafka {
	return &kafka{
		Config: config,
		Logger: logger,
	}
}

func (k *kafka) Start() error {
	config := sarama.NewConfig()

	producer, err := sarama.NewAsyncProducer(k.Config.Brokers, config)
	if err != nil {
		return err
	}

	k.producer = producer
	return nil
}

func (k *kafka) Stop() error {
	err := k.producer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (k *kafka) Write(event model.Event) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: "events",
		Value: sarama.ByteEncoder(bytes),
	}
	k.producer.Input() <- msg
	k.Logger.Info("New message sent")
	return nil
}
