package v1

import (
	"context"
	"crypto/tls"
	"encoding/json"
	v12 "github.com/hoangbktech/cinema-bookings/document-service/pkg/api/v1"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/api/v1"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/kafka"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/mail"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/model"
	"github.com/tkanos/gonfig"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net/smtp"
	"path"
	"path/filepath"
	"runtime"
)

const (
	// apiVersion is version of API is provided by server
	apiVersion = "v1"
	mailConfig = "../../config/mail/config.json"
)

var (
	notificationServer *notificationServiceServer
	ctx context.Context
)

type notificationServiceServer struct {
	documentClient v12.DocumentServiceClient
}



func NewNotificationServiceServer(contxt context.Context, consumer kafka.KafkaConsumer, documentClient v12.DocumentServiceClient) v1.NotificationServiceServer  {
	server := &notificationServiceServer{documentClient:documentClient}
	notificationServer = server
	ctx = contxt
	consumer.Consume(handle)
	return server
}

func handle(data []byte) error {
	notification := &model.Notification{}
	err := json.Unmarshal(data, &notification)
	if err != nil {
		log.Print("unable to parse notification data ", notification)
	}

	switch notification.Method {
		case model.EMAIL:
			notificationServer.SendEmail(notification)
			break
		case model.SMS:
			notificationServer.SendSMS(notification)
			break
	}

	return nil
}

// checkAPI checks if the API version requested by client is supported by server
func (s *notificationServiceServer) checkAPI(api string) error {
	// API version is "" means use current version of the service
	if len(api) > 0 {
		if apiVersion != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements API version '%s', but asked for '%s'", apiVersion, api)
		}
	}
	return nil
}

func (s *notificationServiceServer) SendEmail(notification *model.Notification) (int64, error) {
	configuration := model.MailConfig{}
	err := gonfig.GetConf(getFileName(mailConfig), &configuration)

	ma := mail.Mail{}
	ma.SenderId = configuration.SenderId

	switch notification.Type {
		case model.ALERT :
			htmlResponse, err := notificationServer.documentClient.CreateAlertEmailTemplate(ctx, &v12.AlertTemplateRequest{Api: apiVersion, Cinema: notification.Payload.Cinema, Movie: notification.Payload.Movie})
			if err != nil {
				log.Fatal("there is error while calling document service for template generating")
				return 0, err
			}
			ma.ToIds = []string{configuration.AdminId}
			ma.Subject = configuration.AlertSubject
			ma.Body = htmlResponse.Html
			break
		case model.INFORM:
			break
	}

	messageBody := ma.BuildMessage()

	smtpServer := mail.SmtpServer{Host: configuration.Host, Port: configuration.Port}

	// Build an auth
	auth := smtp.PlainAuth("", ma.SenderId, configuration.Password, smtpServer.Host)

	// Gmail will reject connection if it's not secure
	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer.Host,
	}

	conn, err := tls.Dial("tcp", smtpServer.ServerName(), tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	client, err := smtp.NewClient(conn, smtpServer.Host)
	if err != nil {
		log.Panic(err)
	}

	//Authentication
	if err = client.Auth(auth); err != nil {
		log.Panic(err)
	}

	if err = client.Mail(ma.SenderId); err != nil {
		log.Panic(err)
	}
	for _, k := range ma.ToIds {
		if err = client.Rcpt(k); err != nil {
			log.Panic(err)
		}
	}

	w, err := client.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(messageBody))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	client.Quit()

	return 0, nil
}

func (s *notificationServiceServer) SendSMS(notification *model.Notification) (int64, error) {

	return 0, nil
}

func getFileName(filename string) string {
	_, dirname, _, _ := runtime.Caller(0)
	filePath := path.Join(filepath.Dir(dirname), filename)

	return filePath
}



