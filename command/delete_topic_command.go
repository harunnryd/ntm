package command

import (
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