package user

import (
	"errors"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindByID(id uint) (*User, error)
	Create(user *User) error
}

type gormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) UserRepository {
	return &gormUserRepo{db: db}
}

func (g gormUserRepo) FindByID(id uint) (*User, error) {
	var u User
	if err := g.db.First(&u, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}

		return nil, nil
	}

	return &u, nil
}

func (g gormUserRepo) Create(user *User) error {
	return g.db.Create(user).Error
}
