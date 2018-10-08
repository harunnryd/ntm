package queryApi

import (
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/db"
)

func NewListStatusHandler(ctx echo.Context) error {
	var listStatus []struct {
		ID uuid.UUID `json:"id"`
		Name string `json:"name"`
	}

	db.
		NewORM().
		Connect().
		LogMode(true).
		Raw(`
			select * from statuses
		`).Scan(&listStatus)

	return ctx.JSON(http.StatusOK, listStatus)
}