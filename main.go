package main

import (
	"flag"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"ntm/commandApi"
	"ntm/db"
	"ntm/kafka"
	"ntm/model"
	"ntm/queryApi"
	"time"
)

func main() {
	d := db.NewORM().Connect()
	d.LogMode(true)


	act := flag.String("act", "rest", "Either: rest or consumer")
	flag.Parse()

	fmt.Printf("Welcome to quizes service: %s\n\n", *act)

	switch *act {
	case "rest":
		fmt.Print("Welcome to REST\n")

		e := echo.New()
		status := e.Group("statuses")
		status.GET("", queryApi.NewListStatusHandler)
		status.POST("", commandApi.NewCreateStatusHandler)
		status.PATCH("/:status_id", commandApi.NewUpdateStatusHandler)
		status.DELETE("/:status_id", commandApi.NewDeleteStatusHandler)

		topic := e.Group("topics")
		topic.GET("", queryApi.NewListTopicHandler)
		topic.POST("", commandApi.NewCreateTopicHandler)
		topic.PATCH("/:topic_id", commandApi.NewUpdateTopicHandler)
		topic.DELETE("/:topic_id", commandApi.NewDeleteTopicHandler)

		news := e.Group("news")
		news.GET("", queryApi.NewListNewsHandler)
		news.POST("", commandApi.NewCreateNewsHandler)
		news.PATCH("/:news_id", commandApi.NewUpdateNewsHandler)
		news.DELETE("/:news_id", commandApi.NewDeleteNewsHandler)

		news.POST("/:news_id/topics", commandApi.NewAssignTopicHandler)
		news.PATCH("/:news_id/topics", commandApi.NewUnAssignTopicHandler)
		news.PATCH("/:news_id/statuses", commandApi.NewChangeNewsStatusHandler)

		e.Debug = true
		e.Logger.Fatal(e.Start(":8080"))
	case "consumer":
		fmt.Print("Welcome to CONSUMER\n")
		kafka.NewConsumer("g-ntm", "statuses-topic", "tags-topic", "news-topic").Listen()
	case "testing":
		fmt.Print("Welcome to TESTING\n")
		fmt.Printf("write this you want!")
	case "migrate":
		fmt.Print("Welcome to Migrate\n")

		d.DropTableIfExists(
			new(model.Status),
			new(model.News),
			new(model.Topic)).
			AutoMigrate(
				new(model.Status),
				new(model.News),
				new(model.Topic))

		d.Exec(`insert into statuses (id, name, created_at, updated_at) values (?, ?, ?, ?)`, uuid.Must(uuid.NewV4()), "publish", time.Now(), time.Now())
		d.Exec(`insert into statuses (id, name, created_at, updated_at) values (?, ?, ?, ?)`, uuid.Must(uuid.NewV4()), "draft", time.Now(), time.Now())
		d.Exec(`insert into statuses (id, name, created_at, updated_at) values (?, ?, ?, ?)`, uuid.Must(uuid.NewV4()), "delete", time.Now(), time.Now())
	}
}
