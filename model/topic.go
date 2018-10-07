package model

type Topic struct {
	Schema
	Name string
}

func NewTopic() *Topic {
	return new(Topic)
}