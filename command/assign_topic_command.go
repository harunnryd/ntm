package command

import "github.com/satori/go.uuid"

type AssignTopicCommand struct {
	ID uuid.UUID
	Type string
	NewsID uuid.UUID
	TopicIds []uuid.UUID
}

func NewAssignTopicCommand(newsID uuid.UUID, topicIds... uuid.UUID) *AssignTopicCommand {
	cmd := new(AssignTopicCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "AssignTopicCommand"
	cmd.NewsID = newsID
	cmd.TopicIds = topicIds
	return cmd
}

func (cmd *AssignTopicCommand) Process() error {
	return nil
}
