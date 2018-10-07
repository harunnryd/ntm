package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"os"
)

type ORM struct {
	Dialect string
	Host string
	Port string
	User string
	DatabaseName string
	Password string
	SslMode string
}

func NewORM() *ORM {
	o := new(ORM)
	o.Dialect = os.Getenv("DB_DIALECT")
	o.Host = os.Getenv("DB_HOST")
	o.Port = os.Getenv("DB_PORT")
	o.User = os.Getenv("DB_USER")
	o.DatabaseName = os.Getenv("DB_NAME")
	o.Password = os.Getenv("DB_PASSWORD")
	o.SslMode = os.Getenv("DB_SSLMODE")
	return o
}

func (o ORM) Connect() (db *gorm.DB) {
	db, err := gorm.Open(
		o.Dialect,
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			o.Host,
			o.Port,
			o.User,
			o.DatabaseName,
			o.Password,
			o.SslMode))
	if err != nil {
		fmt.Print("Please wait i'm trying to connect database ...\n")
		os.Exit(-1)
	} else {
		fmt.Printf("Yeay i've been connected database ...\n")
	}

	return
}
