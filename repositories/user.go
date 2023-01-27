package repositories

import (
	"go-commerce/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindById(id string) (entities.User, error)
	FindByEmail(email string) (entities.User, error)
	Create(user entities.User) (entities.User, error)
	Update(user entities.User) (entities.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (u *userRepository) FindById(id string) (entities.User, error) {
	var user entities.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) Create(user entities.User) (entities.User, error) {
	err := u.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userRepository) Update(user entities.User) (entities.User, error) {
	err := u.db.Save(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
