package command

import (
	"github.com/satori/go.uuid"
)

type CreateStatusCommand struct {
	ID uuid.UUID
	Type string
	Name string
}

func NewCreateStatusCommand(name string) *CreateStatusCommand {
	cmd := new(CreateStatusCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "CreateStatusCommand"
	cmd.Name = name
	return cmd
}