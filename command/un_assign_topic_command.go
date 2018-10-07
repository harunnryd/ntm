package command

import "github.com/satori/go.uuid"

type UnAssignTopicCommand struct {
	ID uuid.UUID
	Type string
	NewsID uuid.UUID
	TopicIds []uuid.UUID
}

func NewUnAssignTopicCommand(newsID uuid.UUID, topicIds... uuid.UUID) *UnAssignTopicCommand {
	cmd := new(UnAssignTopicCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "UnAssignTopicCommand"
	cmd.NewsID = newsID
	cmd.TopicIds = topicIds
	return cmd
}

func (cmd *UnAssignTopicCommand) Process() error {
	return nil
}
