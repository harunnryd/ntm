package model

type Status struct {
	Schema
	Name string `gorm:"column:name;"`
}

func NewStatus() *Status {
	return new(Status)
}
