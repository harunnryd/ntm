package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"reflect"
)
type Producer struct {
	Topic string
	Command interface{}
}

func NewProducer(topic string, command interface{}) *Producer {
	p := new(Producer)
	p.Topic = topic
	p.Command = command
	return p
}

func (p *Producer) config() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.ChannelBufferSize = 1
	config.Version = sarama.V0_10_1_0
	return config
}

func (p *Producer) connect() sarama.SyncProducer {
	prod, err := sarama.NewSyncProducer([]string{os.Getenv("KAFKA_HOST")}, p.config())
	if err != nil {
		fmt.Printf("Producer err: %s\n", err)
		os.Exit(-1)
	}
	return prod
}

func (p *Producer) Dispatch() {
	prod := p.connect()
	defer func() {
		if err := prod.Close(); err != nil {
			fmt.Printf("Producer err: %s\n", err)
		}
	}()

	v, _ := json.Marshal(p.Command)

	msg := new(sarama.ProducerMessage)

	msg.Topic = p.Topic
	msg.Key = sarama.StringEncoder(fmt.Sprint(reflect.TypeOf(p.Command).Elem().Name()))
	msg.Value = sarama.StringEncoder(v)

	partition, offset, err := prod.SendMessage(msg)
	if err != nil {
		fmt.Printf("Producer err: %s\n", err)
	} else {
		fmt.Printf("Message: %+v\n", p.Command)
		fmt.Printf("Message stored in key: %s, partition: %d, offset: %d\n", reflect.TypeOf(p.Command).Elem().Name(), partition, offset)
	}
}