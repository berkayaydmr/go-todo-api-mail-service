package common

import (
	"context"
	json "encoding/json"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"mail-service/models"
	"mail-service/services"
	"time"
)

type KafkaEvent struct {
	Network  string
	GroupId  string
	MaxBytes int
	Address  string
	Topic    string
	Partion  int
}

func (k *KafkaEvent) Consumer() {
	config := kafka.ReaderConfig{Topic: k.Topic, GroupID: k.GroupId, Brokers: []string{k.Address}, MaxWait: time.Second}
	reader := kafka.NewReader(config)

	for {
		zap.S().Info(`message`)
		messages, err := reader.ReadMessage(context.Background())
		if err != nil {
			zap.S().Error("Error: ", err)
			continue
		}

		var messageJson models.KafkaUserMail

		if err := json.Unmarshal(messages.Value, &messageJson); err != nil {
			zap.S().Error("Error: ", err)
			continue
		}

		if messageJson.Event == "validateUserMail" {
			err = services.MailService(messageJson)
			if err != nil {
				zap.S().Error("Error:", err)
				continue
			}
		}
		continue
	}
}
