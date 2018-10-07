package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type UpdateStatusHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewUpdateStatusHandler(ctx echo.Context) error {
	h := new(UpdateStatusHandler)

	f := new(struct{
		Name string `json:"name"`
	})

	statusID := uuid.FromStringOrNil(ctx.Param("status_id"))

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewUpdateStatusCommand(statusID, f.Name)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("statuses-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.StatusID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}