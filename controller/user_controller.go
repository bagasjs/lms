package controller

import (
	"github.com/bagasjs/lms/internal/service"
	"github.com/labstack/echo/v4"
)

type UserController struct {
    userService service.UserService
}

func NewUserController(service service.UserService) *UserController {
    return &UserController{
        userService: service,
    }
}

func (controller *UserController) Route(g *echo.Group) {
    g.GET("", controller.AllUsers)
    g.POST("", controller.CreateUser)
    g.GET("/:id", controller.ViewUser)
    g.PUT("/:id", controller.UpdateUser)
    g.DELETE("/:id", controller.DestroyUser)
}
