package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type CreateNewsCommand struct {
	ID uuid.UUID
	Type string
	Title string
	Body string
	TopicIds []uuid.UUID
}

func NewCreateNewsCommand(title, body string, topicIds... uuid.UUID) *CreateNewsCommand {
	cmd := new(CreateNewsCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateNewsCommand"
	cmd.Title = title
	cmd.Body = body
	cmd.TopicIds = topicIds
	return cmd
}

func (cmd *CreateNewsCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}