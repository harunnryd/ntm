package model

import (
	"github.com/satori/go.uuid"
	"time"
)

type Schema struct {
	ID uuid.UUID `gorm:"primary_key;type:char(36);column:id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
