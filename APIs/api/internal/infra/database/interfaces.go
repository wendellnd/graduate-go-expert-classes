package database

import "github.com/wendellnd/graduate-go-expert-classes/api/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
