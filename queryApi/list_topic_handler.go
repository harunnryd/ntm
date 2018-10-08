package queryApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/db"
)

func NewListTopicHandler(ctx echo.Context) error {
	var listTopic []struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
	}

	db.
		NewORM().
		Connect().
		LogMode(true).
		Raw(`
			select * from topics
		`).Scan(&listTopic)

	return ctx.JSON(http.StatusOK, listTopic)
}