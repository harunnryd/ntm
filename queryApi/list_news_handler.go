package queryApi

import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"ntm/db"
)

func NewListNewsHandler(ctx echo.Context) error {
	var (
		listNews []struct {
			ID uuid.UUID `json:"id"`
			Status string `json:"status,omitempty"`
			Topic string `json:"topic,omitempty"`
			Title string `json:"title"`
			Body string `json:"body"`
		}
	)

	status := ctx.QueryParam("status")
	topic := ctx.QueryParam("topic")

	if status != "" && topic != "" {
		fmt.Print("taiii")
		db.
			NewORM().
			Connect().
			LogMode(true).
			Raw(`
			select 
				n.id as id,
				n.title as title,
				n.body as body,
				t.name as topic,
				s.name as status
			from news_topics nt 
			inner join topics t on nt.topic_id = t.id 
			inner join news n on nt.news_id = n.id
			inner join statuses s on n.status_id = s.id
			where s.name = ? and t.name = ?
		`, status, topic).Scan(&listNews)
	} else if status != "" {
		db.
			NewORM().
			Connect().
			LogMode(true).
			Raw(`
			select 
				n.id as id,
				n.title as title,
				n.body as body,
				t.name as topic,
				s.name as status
			from news_topics nt 
			inner join topics t on nt.topic_id = t.id 
			inner join news n on nt.news_id = n.id
			inner join statuses s on n.status_id = s.id
			where s.name = ?
		`, status).Scan(&listNews)
	} else if topic != "" {
		db.
			NewORM().
			Connect().
			LogMode(true).
			Raw(`
			select 
				n.id as id,
				n.title as title,
				n.body as body,
				t.name as topic,
				s.name as status
			from news_topics nt 
			inner join topics t on nt.topic_id = t.id 
			inner join news n on nt.news_id = n.id
			inner join statuses s on n.status_id = s.id
			where t.name = ?
		`, topic).Scan(&listNews)
	} else {
		db.
			NewORM().
			Connect().
			LogMode(true).
			Raw(`
			select * from news
		`).Scan(&listNews)
	}


	return ctx.JSON(http.StatusOK, listNews)
}