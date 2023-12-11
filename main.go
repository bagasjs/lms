package main

import (
	"database/sql"
	"log"
	"os"
	"path"

	"github.com/bagasjs/lms/controller"
	"github.com/bagasjs/lms/internal/repository"
	"github.com/bagasjs/lms/internal/service"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"github.com/spf13/cobra"
)

func Serve() {
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

func RunMigrations(filePath string) {
    dir, err := os.Open(filePath)
    if err != nil {
        log.Fatalln("Failed to open directory ", filePath, ": ", err)
        return
    }

    files, err := dir.ReadDir(0)
    if err != nil {
        log.Fatalln("Failed to read directory ", filePath, ": ", err);
        return
    }

	db, err := sql.Open("sqlite3", "./db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
    for _, f := range files {
        fileContent, err := os.ReadFile(path.Join(dir.Name(), f.Name()))
        if err != nil {
            log.Println("Failed to open file ", path.Join(dir.Name(), f.Name()), ": ", err)
        }

        _, err = db.Exec(string(fileContent))
        if err != nil {
            log.Fatalln("ERROR on file ", f.Name(), ":", err)
        }
    }
}

func main() {

    cmd := &cobra.Command {
        Use: "lms",
        Short: "LMS is a simple Learning Management System by @bagasjs",
        Run: func(cmd *cobra.Command, args []string) {
        },
    }

    serve := &cobra.Command {
        Use: "serve",
        Short: "Run the server",
        Run: func(cmd *cobra.Command, args []string) {
            Serve()
        },
    }


    migrateUp := &cobra.Command {
        Use: "migrate:up",
        Short: "Run the migration at migrations/",
        Run: func(cmd *cobra.Command, args []string) {
            RunMigrations("./migrations/up/")
        },
    }

    migrateDown := &cobra.Command {
        Use: "migrate:down",
        Short: "Run the migration at migrations/",
        Run: func(cmd *cobra.Command, args []string) {
            RunMigrations("./migrations/down/")
        },
    }

    cmd.AddCommand(serve)
    cmd.AddCommand(migrateUp);
    cmd.AddCommand(migrateDown);
    cmd.Execute()
}
