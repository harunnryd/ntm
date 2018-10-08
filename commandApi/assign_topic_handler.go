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

type AssignTopicHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	} `json:"promise"`
}

func NewAssignTopicHandler(ctx echo.Context) error {
	h := new(CreateNewsHandler)

	f := new(struct{
		TopicID uuid.UUID `json:"topic_id" valid:"required"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(f)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	newsID := uuid.FromStringOrNil(ctx.Param("news_id"))

	cmd := command.NewAssignTopicCommand(newsID, f.TopicID)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("news-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.NewsID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}