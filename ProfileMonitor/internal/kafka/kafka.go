package kafka

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/JohnnyJa/AdServer/ProfileMonitor/service"
	"github.com/sirupsen/logrus"
)

type ProfileForKafka struct {
	ID        string             `json:"id"`
	Name      string             `json:"name"`
	Creatives []CreativeForKafka `json:"creatives"`
}

type CreativeForKafka struct {
	ID           string            `json:"id"`
	MediaURL     string            `json:"media_url"`
	Width        int               `json:"width"`
	Height       int               `json:"height"`
	CreativeType string            `json:"creative_type"`
	Targeting    map[string]string `json:"targeting"`
}

type Kafka interface {
	service.Service
	Write(forKafka []ProfileForKafka) error
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

func (k *kafka) Write(forKafka []ProfileForKafka) error {
	bytes, err := json.Marshal(forKafka)
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
