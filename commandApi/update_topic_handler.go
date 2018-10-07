package commandApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/kafka"
	"ntm/command"
	"os"
)

type UpdateTopicHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	}
}

func NewUpdateTopicHandler(ctx echo.Context) error {
	h := new(UpdateTopicHandler)

	f := new(struct{
		Name string `json:"name"`
	})

	topicID := uuid.FromStringOrNil(ctx.Param("topic_id"))

	if err := ctx.Bind(f); err != nil {
		return err
	}

	cmd := command.NewUpdateTopicCommand(topicID, f.Name)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("tags-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.TopicID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusAccepted, h)
}