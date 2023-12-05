package main

import (
	"database/sql"
	"log"

	"github.com/bagasjs/lms/controller"
	"github.com/bagasjs/lms/internal/repository"
	"github.com/bagasjs/lms/internal/service"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    app := echo.New()

    userRepository := repository.NewUserSQLite3Repository(db)
    userController := controller.NewUserController(service.NewUserService(userRepository))
    userController.Route(app.Group("/api/users"))
    app.Logger.Fatal(app.Start(":6969"))
}
