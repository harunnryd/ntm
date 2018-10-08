package commandApi

import (
	"github.com/asaskevich/govalidator"
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
	} `json:"promise"`
}

func NewCreateTopicHandler(ctx echo.Context) error {
	h := new(CreateTopicHandler)

	f := new(struct{
		Name string `json:"name" valid:"required"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(f)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
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