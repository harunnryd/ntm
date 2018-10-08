package command

import "github.com/satori/go.uuid"

type ChangeNewsStatusCommand struct {
	ID uuid.UUID
	Type string
	NewsID uuid.UUID
	StatusID uuid.UUID
}

func NewChangeNewsStatusCommand(newsID uuid.UUID, statusID uuid.UUID) *ChangeNewsStatusCommand {
	cmd := new(ChangeNewsStatusCommand)
	cmd.ID = uuid.Must(uuid.NewV4())
	cmd.Type = "ChangeNewsStatusCommand"
	cmd.NewsID = newsID
	cmd.StatusID = statusID
	return cmd
}