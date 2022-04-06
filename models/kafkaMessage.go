package models

type KafkaUserMail struct {
	Event        string `json:"event"`
	MailReceiver string `json:"mailReceiver"`
	UserID       string `json:"userID"`
}
