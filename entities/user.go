package entities

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Base
}

func (u *User) BeforeCreate(db *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	if u.ID == uuid.Nil {
		err = errors.New("ID is empty")
	}
	return
}
