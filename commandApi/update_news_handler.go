package commandApi

import (
	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/command"
	"ntm/kafka"
	"os"
)

type UpdateNewsHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	} `json:"promise"`
}

func NewUpdateNewsHandler(ctx echo.Context) error {
	h := new(UpdateNewsHandler)

	f := new(struct{
		Title string `json:"title" valid:"required"`
		Body string `json:"body" valid:"required"`
	})

	newsID := uuid.FromStringOrNil(ctx.Param("news_id"))

	if err := ctx.Bind(f); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(f)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
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