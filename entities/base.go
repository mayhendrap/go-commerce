package entities

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
	"time"
)

type Base struct {
	gorm.Model
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
	CreatedBy uuid.UUID
	UpdatedBy uuid.UUID
	DeletedBy uuid.UUID
}
