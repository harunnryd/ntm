package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type DeleteNewsHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	} `json:"promise"`
}

func NewDeleteNewsHandler(ctx echo.Context) error {
	h := new(DeleteNewsHandler)

	newsID := uuid.FromStringOrNil(ctx.Param("news_id"))

	cmd := command.NewDeleteNewsCommand(newsID)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("statuses-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.NewsID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}