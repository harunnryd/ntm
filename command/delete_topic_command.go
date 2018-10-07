package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type DeleteTopicCommand struct {
	Type string
	TopicID uuid.UUID
}

func NewDeleteTopicCommand(topicID uuid.UUID) *DeleteTopicCommand {
	cmd := new(DeleteTopicCommand)
	cmd.Type = "DeleteTopicCommand"
	cmd.TopicID = topicID
	return cmd
}

func (cmd *DeleteTopicCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}