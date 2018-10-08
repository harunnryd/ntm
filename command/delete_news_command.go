package command

import (
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