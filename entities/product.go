package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Product struct {
	Base
	Title       string
	Description string
	Price       int
}

func (p *Product) BeforeCreate(db *gorm.DB) (err error) {
	p.ID = uuid.NewV4()
	if p.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
