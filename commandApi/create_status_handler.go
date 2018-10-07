package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type CreateStatusHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewCreateStatusHandler(ctx echo.Context) error {
	h := new(CreateStatusHandler)

	f := new(struct{
		Name string `json:"name"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewCreateStatusCommand(f.Name)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("statuses-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.ID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusCreated, h)
}