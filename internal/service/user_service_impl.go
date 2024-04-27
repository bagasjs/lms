package service

import (
	"errors"
	"strings"

	"github.com/bagasjs/lms/internal/entity"
	"github.com/bagasjs/lms/internal/model"
	"github.com/bagasjs/lms/internal/repository"
)

type userServiceImpl struct {
	UserRepository repository.UserRepository
}

func (service *userServiceImpl) Create(request model.CreateUserRequest) (response model.GenericUserResponse, err error) {
    if strings.Compare(request.Password, request.PasswordConfirmation) != 0 {
        return response, errors.New("Password is not confirmed")
    }
    user := entity.User{ Name: request.Name, Password: request.Password, Email: request.Email }
    err = service.UserRepository.Insert(user)
    if err != nil {
        return response, err
    }
    response.Name = user.Name
    response.Email = user.Email
    response.ID = 0
    return response, err
}

func (serv *userServiceImpl) Update(request model.UpdateUserRequest) (response model.GenericUserResponse, err error) {
    user := entity.User{ 
        ID: request.ID, 
        Name: request.Name, 
        Password: request.Password, 
        Email: request.Email,
    }
    err = serv.UserRepository.Update(user)
    if err != nil {
        return response, err
    }
    response.Name = user.Name
    response.Email = user.Email
    response.ID = user.ID
    return response, err
}

func (service *userServiceImpl) List() (responses []model.GenericUserResponse, err error) {
	users, err := service.UserRepository.FindAll()
	if err != nil {
		return nil,  err
	}

	for _, user := range users {
		responses = append(responses, model.GenericUserResponse {
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		})
	}

    return responses, err
}

func (service *userServiceImpl) FindByID(id string) (response model.GenericUserResponse, err error) {
    user, err := service.UserRepository.FindByID(id)
    if err != nil {
        return response, err
    }
    response.ID = user.ID
    response.Name = user.Name
    response.Email = user.Email
    return response, err
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userServiceImpl{
		UserRepository: userRepository,
	}
}


