package command

import "github.com/satori/go.uuid"

type AssignTopicCommand struct {
	Type string
	NewsID uuid.UUID
	TopicID uuid.UUID
}

func NewAssignTopicCommand(newsID uuid.UUID, topicID uuid.UUID) *AssignTopicCommand {
	cmd := new(AssignTopicCommand)
	cmd.Type = "AssignTopicCommand"
	cmd.NewsID = newsID
	cmd.TopicID = topicID
	return cmd
}
