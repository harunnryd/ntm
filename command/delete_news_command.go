package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type DeleteNewsCommand struct {
	Type string
	NewsID uuid.UUID
}

func NewDeleteNewsCommand(newsID uuid.UUID) *DeleteNewsCommand {
	cmd := new(DeleteNewsCommand)
	cmd.Type = "DeleteNewsCommand"
	cmd.NewsID = newsID
	return cmd
}

func (cmd *DeleteNewsCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}