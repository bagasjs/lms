package repository

import (
	"database/sql"
	"log"

	"github.com/bagasjs/lms/internal/entity"
)

type userRepositoryImplSQLite3 struct {
    db *sql.DB
}

func (repo *userRepositoryImplSQLite3) Insert(user entity.User) error {
	stmt, err := repo.db.Prepare("INSERT INTO users(email, name, password) values(?, ?, ?)")
    if err != nil {
        return err
    }
	defer stmt.Close()
    _, err = stmt.Exec(user.Email, user.Name, user.Password)
    if err != nil {
        return err
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) Update(user entity.User) error {
	stmt, err := repo.db.Prepare("UPDATE users SET email=?, name=?, password=? WHERE id=?")
    if err != nil {
        return err
    }
	defer stmt.Close()
    _, err = stmt.Exec(user.Email, user.Name, user.Password, user.ID)
    if err != nil {
        return err
    }
    return nil
}

func (repo *userRepositoryImplSQLite3) FindAll() (users []entity.User, err error) {
    rows, err := repo.db.Query("SELECT * FROM users")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    user := entity.User{}
    for rows.Next() {
        err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
        if err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    return users, nil
}

func (repo *userRepositoryImplSQLite3) DeleteAll() error {
    return nil
}

func (repo *userRepositoryImplSQLite3) FindByID(id string) (user entity.User, err error) {
    rows, err := repo.db.Query("SELECT * FROM users WHERE id=?", id)
    if err != nil {
        return user, err
    }
    defer rows.Close()
    rows.Next()
    err = rows.Scan(&user.ID, &user.Email, &user.Name, &user.Password)
    if err != nil {
        return user, err
    }

    return user, nil
}

func NewUserSQLite3Repository(db *sql.DB) UserRepository {
    _, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        email VARCHAR(255) NOT NULL UNIQUE,
        name VARCHAR(255) NOT NULL,
        password VARCHAR(255) NOT NULL
    )`)

    if err != nil {
        log.Fatal(err)
    }

    return &userRepositoryImplSQLite3{
        db: db,
    }
}
