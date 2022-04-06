package mails

import (
	"bytes"
	"fmt"
	"html/template"
	"mail-service/models"
)

func MailParse(messageJson models.KafkaUserMail) ([]byte, error) {

	tmpl, err := template.ParseFiles("/Users/berkayaydemir/GolandProjects/mail-service/mails/ValidationMail.html")
	if err != nil {
		return nil, err
	}

	data := models.KafkaUserMail{MailReceiver: messageJson.MailReceiver}

	buffer := new(bytes.Buffer)
	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	buffer.Write([]byte(fmt.Sprintf("Subject: Validate your email \n%s\n\n", mimeHeaders)))

	if err = tmpl.Execute(buffer, data); err != nil {
		return nil, err
	}

	message := []byte(buffer.String())

	return message, err
}
