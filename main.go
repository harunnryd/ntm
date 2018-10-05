package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"time"
)

var (
	DB = initDB()
)

type Schema struct {
	ID uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type News struct {
	Schema
	Title string

}

func initDB() (db *gorm.DB) {
	db, err := gorm.Open(
		os.Getenv("DB_DIALECT"),
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_SSLMODE")))
	if err != nil {
		fmt.Print("Please wait i'm trying to connect database ...\n")
		os.Exit(-1)
	} else {
		fmt.Printf("Yeay i've been connected database ...\n")
	}

	return
}

func main() {
	DB.LogMode(true)
	DB.DropTableIfExists(new(News)).AutoMigrate(new(News))
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Debug = true
	e.Logger.Fatal(e.Start(":8080"))
}
