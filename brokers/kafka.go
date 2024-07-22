package brokers

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"testServ/models"
)

var producer sarama.SyncProducer

func InitKafka(brokers []string) error {
	var err error
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err = sarama.NewSyncProducer(brokers, config)
	return err
}

func ProduceMessage(message models.Message) error {
	msg, err := json.Marshal(message)
	if err != nil {
		return err
	}

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "messages",
		Value: sarama.StringEncoder(msg),
	})

	return err
}
