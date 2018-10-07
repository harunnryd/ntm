package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type DeleteTopicHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewDeleteTopicHandler(ctx echo.Context) error {
	h := new(DeleteTopicHandler)

	topicID := uuid.FromStringOrNil(ctx.Param("topic_id"))

	cmd := command.NewDeleteTopicCommand(topicID)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("tags-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.TopicID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}