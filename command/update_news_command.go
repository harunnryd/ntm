package command

import (
	"fmt"
	"github.com/satori/go.uuid"
)

type UpdateNewsCommand struct {
	Type string
	NewsID uuid.UUID
	Title string
	Body string
}

func NewUpdateNewsCommand(newsID uuid.UUID,  title, body string) *UpdateNewsCommand {
	cmd := new(UpdateNewsCommand)
	cmd.Type = "UpdateNewsCommand"
	cmd.NewsID = newsID
	cmd.Title = title
	cmd.Body = body
	return cmd
}

func (cmd *UpdateNewsCommand) Process() error {
	fmt.Printf("process: %+v\n", cmd)
	return nil
}