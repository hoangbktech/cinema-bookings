package kafka

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
	SynProducer sarama.SyncProducer
}

func (kp *KafkaProducer) Init() {
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
	if err != nil {
		log.Fatal("there is an error in initializing Producer")
	}
	kp.SynProducer = prd
}

func (kp *KafkaProducer) Publish(message *model.Notification) {

	data, err := json.Marshal(message)

	// publish sync
	msg := &sarama.ProducerMessage {
		Topic: kp.Topic,
		Value: sarama.ByteEncoder(data),
	}
	par, o, err := kp.SynProducer.SendMessage(msg)
	if err != nil {
		fmt.Println("Error publish: ", err.Error())
	}

	// publish async
	//producer.Input() <- &sarama.ProducerMessage{

	fmt.Println("Partition: ", par)
	fmt.Println("Offset: ", o)
}

