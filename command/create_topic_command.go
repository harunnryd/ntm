package command

import "github.com/satori/go.uuid"

type CreateTopicCommand struct {
	ID uuid.UUID
	Type string
	Name string
}

func NewCreateTopicCommand(name string) *CreateTopicCommand {
	cmd := new(CreateTopicCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateTopicCommand"
	cmd.Name = name
	return cmd
}