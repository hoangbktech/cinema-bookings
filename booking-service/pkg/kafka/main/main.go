package main

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hoangbktech/cinema-bookings/booking-service/pkg/model"
	"log"
	"os"
)
var (
	maxRetry = 5
)

type KafkaProducer struct {
	KafkaHost string
	Topic string
}

func (kp *KafkaProducer) InitProducer()(sarama.SyncProducer, error) {
	// setup sarama log to stdout
	sarama.Logger = log.New(os.Stdout, "", log.Ltime)

	// producer config
	config := sarama.NewConfig()
	config.Producer.Retry.Max = maxRetry
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true

	// async producer
	//prd, err := sarama.NewAsyncProducer([]string{kafkaConn}, config)

	// sync producer
	prd, err := sarama.NewSyncProducer([]string{kp.KafkaHost}, config)

	return prd, err
}

func (kp *KafkaProducer) Publish(message *model.Notification) {

	producer, err := kp.InitProducer()
	if err != nil {
		fmt.Println("Error producer: ", err.Error())
		os.Exit(1)
	}

	data, err := json.Marshal(message)

	// publish sync
	msg := &sarama.ProducerMessage {
		Topic: kp.Topic,
		Value: sarama.ByteEncoder(data),
	}
	par, o, err := producer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", par)
	fmt.Println("Offset: ", o)
}

func main() {
	notification := &model.Notification{
		Type: model.ALERT,
		Method: model.EMAIL,
		Payload: model.Payload{
			Movie:          "movie1",
			Cinema:         "Cinema1",
			TotalSeats:     100,
			AvailableSeats: 50,
			BookingUser: model.User{
				Name:        "Test",
				LastName:    "Test1",
				Email:       "hoangbktech@gmail.com",
				PhoneNumber: "0934347073",
			},
		},
	}


	kafkaProducer := KafkaProducer{KafkaHost: "localhost:9092", Topic: "notification"}
	kafkaProducer.Publish(notification)
}
