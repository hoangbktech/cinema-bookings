package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/hoangbktech/cinema-bookings/notification-service/pkg/model"
	"github.com/wvanbergen/kafka/consumergroup"
	"log"
	"time"
)

const (
	cgroup = "zgroup"
)

type KafkaConsumer struct {
	ZookeeperHost string
	Topic string
	ConsumerGroup *consumergroup.ConsumerGroup
}

func (kc *KafkaConsumer) Init(){
	// consumer config
	config := consumergroup.NewConfig()
	config.Offsets.Initial = sarama.OffsetOldest
	config.Offsets.ProcessingTimeout = 10 * time.Second

	// join to consumer group
	cg, err := consumergroup.JoinConsumerGroup(cgroup, []string{kc.Topic}, []string{kc.ZookeeperHost}, config)
	if err != nil {
		log.Fatal("there is an error in consumer initialization")
	}
	kc.ConsumerGroup = cg
}

func (kc *KafkaConsumer) Consume(fn model.ProcessMessage) {
	for {
		select {
		case msg := <-kc.ConsumerGroup.Messages():
			// messages coming through chanel
			// only take messages from subscribed topic
			if msg.Topic != kc.Topic {
				continue
			}

			fmt.Println("Topic: ", msg.Topic)
			fmt.Println("Value: ", string(msg.Value))

			err := fn(msg.Value)
			if err != nil {
				log.Fatal(err)
			}

			// commit to zookeeper that message is read
			// this prevent read message multiple times after restart
			err = kc.ConsumerGroup.CommitUpto(msg)
			if err != nil {
				fmt.Println("Error commit zookeeper: ", err.Error())
			}
		}
	}
}
