package command

import "github.com/satori/go.uuid"

type UnAssignTopicCommand struct {
	Type string
	NewsID uuid.UUID
	TopicID uuid.UUID
}

func NewUnAssignTopicCommand(newsID uuid.UUID, topicID uuid.UUID) *UnAssignTopicCommand {
	cmd := new(UnAssignTopicCommand)
	cmd.Type = "UnAssignTopicCommand"
	cmd.NewsID = newsID
	cmd.TopicID = topicID
	return cmd
}
