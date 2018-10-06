package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type CreateTopicHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewCreateTopicHandler(ctx echo.Context) error {
	h := new(CreateTopicHandler)

	f := new(struct{
		Name string `json:"name"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateTopicCommand(f.Name)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("tags-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.ID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusCreated, h)
}