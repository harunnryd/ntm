package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type DeleteStatusCommand struct {
	Type string
	StatusID uuid.UUID
}

func NewDeleteStatusCommand(statusID uuid.UUID) *DeleteStatusCommand {
	cmd := new(DeleteStatusCommand)
	cmd.Type = "DeleteStatusCommand"
	cmd.StatusID = statusID
	return cmd
}

func (cmd *DeleteStatusCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}