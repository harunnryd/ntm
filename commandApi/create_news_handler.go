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

type CreateNewsHandler struct {
	Promise struct {
		ID uuid.UUID `json:"id"`
		Message string `json:"message"`
	} `json:"promise"`
}

func NewCreateNewsHandler(ctx echo.Context) error {
	h := new(CreateNewsHandler)

	f := new(struct{
		Title string `json:"title" valid:"required"`
		Body string `json:"body" valid:"required"`
	})

	if err := ctx.Bind(f); err != nil {
		return err
	}

	_, err := govalidator.ValidateStruct(f)

	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}


	cmd := command.NewCreateNewsCommand(f.Title, f.Body)

	if os.Getenv("ENV") != "testing" {
		kafka.
			NewProducer("news-topic", cmd).
			Dispatch()
	}

	h.Promise.ID = cmd.ID
	h.Promise.Message = "success!"

	return ctx.JSON(http.StatusCreated, h)
}