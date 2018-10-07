package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type UpdateNewsHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewUpdateNewsHandler(ctx echo.Context) error {
	h := new(UpdateNewsHandler)

	f := new(struct{
		Title string `json:"title"`
		Body string `json:"body"`
	})

	newsID := uuid.FromStringOrNil(ctx.Param("news_id"))

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewUpdateNewsCommand(newsID, f.Title, f.Body)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("news-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.NewsID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}