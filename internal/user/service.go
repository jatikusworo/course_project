package user

import "time"

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
	return s.repo.FindByID(id)
}

func (s service) CreateUser(name, email string) (*User, error) {
	u := &User{Name: name, Email: email, CreatedAt: time.Now()}
	if err := s.repo.Create(u); err != nil {
		return nil, err
	}
	return u, nil
}
