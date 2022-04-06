package envoriment

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"os"
	"strconv"
)

type Envoriment struct {
	From         string
	MailPassword string
	SmtpHost     string
	SmtpPort     string
	Network      string
	GroupId      string
	Address      string
	Topic        string
	Partion      int
}

func GetEnvoriment() *Envoriment {
	err := godotenv.Load(".env")
	if err != nil {
		zap.S().Error("Error ", err)
	}

	from := os.Getenv("FROM")
	mailPassword := os.Getenv("MAILPASSWORD")
	smtpHost := os.Getenv("SMTPHOST")
	smtpPort := os.Getenv("SMTPPORT")
	network := os.Getenv("NETWORK")
	GroupId := os.Getenv("GROUPID")
	addressHost := os.Getenv("ADDRESSHOST")
	addressPort := os.Getenv("ADDRESSPORT")
	topic := os.Getenv("TOPIC")
	partionS := os.Getenv("PARTION")

	partion, _ := strconv.ParseInt(partionS, 10, 64)

	return &Envoriment{
		From:         from,
		MailPassword: mailPassword,
		SmtpHost:     smtpHost,
		SmtpPort:     smtpPort,
		Network:      network,
		GroupId:      GroupId,
		Address:      addressHost + ":" + addressPort,
		Topic:        topic,
		Partion:      int(partion),
	}
}
