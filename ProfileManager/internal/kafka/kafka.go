package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/JohnnyJa/AdServer/ProfileManager/internal/model"
	"github.com/JohnnyJa/AdServer/ProfileManager/service"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type Kafka interface {
	service.Service
	Read() (map[uuid.UUID]model.Profile, error)
}

type kafka struct {
	Config   *Config
	Logger   *logrus.Logger
	consumer sarama.Consumer
}

func New(config *Config, logger *logrus.Logger) Kafka {
	return &kafka{
		Config: config,
		Logger: logger,
	}
}

func (k *kafka) Start() error {
	config := sarama.NewConfig()

	consumer, err := sarama.NewConsumer(k.Config.Brokers, config)
	if err != nil {
		return err
	}

	k.consumer = consumer
	return nil
}

func (k *kafka) Stop() error {
	err := k.consumer.Close()
	if err != nil {
		return err
	}
	return nil
}

func (k *kafka) Read() (map[uuid.UUID]model.Profile, error) {

	partitions, err := k.consumer.ConsumePartition(k.Config.Topic, k.Config.Partition, sarama.OffsetOldest)
	if err != nil {
		return nil, err
	}

}

func (k *kafka) Write(event model.Event) error {
	bytes, err := json.Marshal(event)
	if err != nil {
		return err
	}

	msg := &sarama.ProducerMessage{
		Topic: k.Config.Topic,
		Value: sarama.ByteEncoder(bytes),
	}
	k.producer.Input() <- msg
	k.Logger.Info("New message sent")
	return nil
}
