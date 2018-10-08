package model

import "github.com/satori/go.uuid"

type News struct {
	Schema
	Status Status `gorm:"foreignkey:StatusID;"`
	StatusID uuid.UUID `gorm:"index;type:char(36);column:status_id;"`
	Title string `gorm:"column:title"`
	Body string `gorm:"column:body;"`
	Topics []Topic `gorm:"many2many:news_topics;"`
}

func NewNews() *News {
	return new(News)
}