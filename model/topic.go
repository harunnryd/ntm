package model

type Topic struct {
	Schema
	Name string `gorm:"column:name;"`
	News []News `gorm:"many2many:news_topics;"`
}

func NewTopic() *Topic {
	return new(Topic)
}