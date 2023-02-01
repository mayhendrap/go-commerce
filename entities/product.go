package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID          uuid.UUID
	Title       string
	Description string
	Price       int
	Base
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4()
	if p.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
