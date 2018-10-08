package command

import (
	"github.com/satori/go.uuid"
)

type UpdateTopicCommand struct {
	Type string
	TopicID uuid.UUID
	Name string
}

func NewUpdateTopicCommand(topicID uuid.UUID,  name string) *UpdateTopicCommand {
	cmd := new(UpdateTopicCommand)
	cmd.Type = "UpdateTopicCommand"
	cmd.TopicID = topicID
	cmd.Name = name
	return cmd
}