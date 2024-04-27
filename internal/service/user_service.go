package service

import "github.com/bagasjs/lms/internal/model"

type UserService interface {
    Create(request model.CreateUserRequest) (response model.GenericUserResponse, error error)
    List() (response []model.GenericUserResponse, error error)
    FindByID(id string) (response model.GenericUserResponse, error error)
    Update(request model.UpdateUserRequest) (response model.GenericUserResponse, error error)
}
