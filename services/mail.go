package services

import (
	"mail-service/common/envoriment"
	"mail-service/mails"
	"mail-service/models"
	"mail-service/utils"
	"net/smtp"

	"go.uber.org/zap"
)

func MailService(messageJson models.KafkaUserMail) error {
	env := envoriment.GetEnvoriment()

	mail, err := utils.DecodeHash(messageJson.MailReceiver)
	if err != nil {
		zap.S().Error("Error: ", err)
		return err
	}

	to := []string{mail}

	message, err := mails.MailParse(messageJson)
	if err != nil {
		zap.S().Error("Error: ", err)
		return err
	}

	auth := smtp.PlainAuth("", env.From, env.MailPassword, env.SmtpHost)

	err = smtp.SendMail(env.SmtpHost+":"+env.SmtpPort, auth, env.From, to, message)
	if err != nil {
		zap.S().Error("Error: ", err)
		return err
	}
	zap.S().Info("Validation mail sent")

	return err
}
