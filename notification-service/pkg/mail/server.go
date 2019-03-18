package mail

import (
	"fmt"
	"strings"
)

type Mail struct {
	SenderId string
	ToIds    []string
	Subject  string
	Body     string
}

type SmtpServer struct {
	Host string
	Port string
}

func (s *SmtpServer) ServerName() string {
	return s.Host + ":" + s.Port
}

func (mail *Mail) BuildMessage() string {
	message := ""
	message += fmt.Sprintf("From: %s\r\n", mail.SenderId)
	if len(mail.ToIds) > 0 {
		message += fmt.Sprintf("To: %s\r\n", strings.Join(mail.ToIds, ";"))
	}

	message += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	message += "\r\n" + mail.Body

	return message
}
