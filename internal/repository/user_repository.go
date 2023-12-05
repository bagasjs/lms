package repository

import "github.com/bagasjs/lms/internal/entity"

type UserRepository interface {
    Insert(user entity.User) error
    FindAll() (users []entity.User, err error)
    FindByID(id string) (user entity.User, err error)
    DeleteAll() error
}
