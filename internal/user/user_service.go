package user

import (
	"course_project/internal/common"
	"fmt"
	"time"
)

type Service interface {
	GetUser(id uint) (*User, error)
	CreateUser(name, email string) (*User, error)
}

type service struct {
	repo UserRepository
}

func NewService(r UserRepository) Service {
	return &service{repo: r}
}

func (s service) GetUser(id uint) (*User, error) {
	u, err := s.repo.FindByID(id)
	if err != nil {
		return nil, common.NewInternalError(err.Error())
	}

	if u == nil {
		return nil, common.NewNotFound(fmt.Sprintf("user %v does not exist", id))
	}

	return u, nil
}

func (s service) CreateUser(name, email string) (*User, error) {
	u := &User{Name: name, Email: email, CreatedAt: time.Now()}
	if err := s.repo.Create(u); err != nil {
		return nil, common.NewInternalError(err.Error())
	}
	return u, nil
}
