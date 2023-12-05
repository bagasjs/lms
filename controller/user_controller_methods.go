package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/bagasjs/lms/internal/model"
)

func (controller *UserController) AllUsers(c echo.Context) error {
    data, err := controller.userService.List()
    if err != nil {
        errorMsg := fmt.Sprintf("Failed to fetch all users due to \"%s\"", err)
        return c.JSON(http.StatusBadRequest, model.ResponseNotFound(errorMsg))
    }
    return c.JSON(http.StatusOK, model.ResponseOk(data, "User fetched"))
}

func (controller *UserController) ViewUser(c echo.Context) error {
    id := c.Param("id");
    data, err := controller.userService.FindByID(id);
    if err != nil {
        errorMsg := fmt.Sprintf("Failed to fetch user due to \"%s\"", err)
        return c.JSON(http.StatusBadRequest, model.ResponseNotFound(errorMsg))
    }
    return c.JSON(http.StatusOK, model.ResponseOk(data, "User fetched"))
}

func (controller *UserController) CreateUser(c echo.Context) error {
    createUserRequest := model.CreateUserRequest{}
    if err := c.Bind(&createUserRequest); err != nil {
        return c.JSON(http.StatusBadRequest, model.ResponseBadRequest("Invalid form body"))
    }
    createUserResponse, err := controller.userService.Create(createUserRequest)

    if err != nil {
        errorMsg := fmt.Sprintf("Failed to create user due to \"%s\"", err)
        return c.JSON(http.StatusBadRequest, model.ResponseBadRequest(errorMsg))
    }

    return c.JSON(http.StatusOK, model.ResponseOk(createUserResponse, "User created"))
}
