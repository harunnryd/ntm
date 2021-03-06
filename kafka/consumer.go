package kafka

import (
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"golang.org/x/net/html/atom"
	. "ntm/command"
	"os"
)

type Consumer struct {
	GroupID string
	Topics []string
}

func NewConsumer(groupID string, topics ...string) *Consumer {
	c := new(Consumer)
	c.GroupID = groupID
	c.Topics = topics
	return c
}

func (c *Consumer) config() *cluster.Config {
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Group.Return.Notifications = true
	return config
}

func (c *Consumer) connect(groupID string, topics ...string) *cluster.Consumer {
	consumer, err := cluster.NewConsumer([]string{os.Getenv("KAFKA_HOST")}, groupID, topics, c.config())
	if err != nil {
		fmt.Printf("Consumer err: %s\n", err)
		os.Exit(-1)
	}
	return consumer
}

func (c *Consumer) Listen() {
	consumer := c.connect(c.GroupID, c.Topics...)
	defer func() {
		if err := consumer.Close(); err != nil {
			fmt.Printf("Consumer err: %s\n", err)
		}
	}()

	go func() {
		for {
			select {
			case err, more := <-consumer.Errors():
				if more {
					fmt.Printf("Consumer err: %s\n", err)
				}
			case ntf, more := <-consumer.Notifications():
				if more {
					fmt.Printf("Rebalanced: %+v\n", ntf)
				}
				
			}
		}
	}()

	for {
		select {
		case msg := <-consumer.Messages():
			key := atom.String(msg.Key)
			consumer.MarkOffset(msg, "")

			switch key {
			case "CreateStatusCommand":
				cmd := new(CreateStatusCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
					fmt.Printf("Process: %+v\n", cmd)
				}
			case "UpdateStatusCommand":
				cmd := new(UpdateStatusCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "DeleteStatusCommand":
				cmd := new(DeleteStatusCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "CreateNewsCommand":
				cmd := new(CreateNewsCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "UpdateNewsCommand":
				cmd := new(UpdateNewsCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "DeleteNewsCommand":
				cmd := new(DeleteNewsCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "CreateTopicCommand":
				cmd := new(CreateTopicCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "UpdateTopicCommand":
				cmd := new(UpdateTopicCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "DeleteTopicCommand":
				cmd := new(DeleteTopicCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "AssignTopicCommand":
				cmd := new(AssignTopicCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "UnAssignTopicCommand":
				cmd := new(UnAssignTopicCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			case "ChangeNewsStatusCommand":
				cmd := new(ChangeNewsStatusCommand)
				if err := json.Unmarshal(msg.Value, cmd); err == nil {
					cmd.Process()
				}
			}
		}
	}
}