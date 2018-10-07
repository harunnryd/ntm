package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type UpdateStatusCommand struct {
	Type string
	StatusID uuid.UUID
	Name string
}

func NewUpdateStatusCommand(statusID uuid.UUID,  name string) *UpdateStatusCommand {
	cmd := new(UpdateStatusCommand)
	cmd.Type = "UpdateStatusCommand"
	cmd.StatusID = statusID
	cmd.Name = name
	return cmd
}

func (cmd *UpdateStatusCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}