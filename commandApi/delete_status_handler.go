package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type DeleteStatusHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	} `json:"promise"`
}

func NewDeleteStatusHandler(ctx echo.Context) error {
	h := new(DeleteStatusHandler)

	statusID := uuid.FromStringOrNil(ctx.Param("status_id"))

	cmd := command.NewDeleteStatusCommand(statusID)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("statuses-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.StatusID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}