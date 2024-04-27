package model

type CreateUserRequest struct {
    Email                string `form:"email"`
    Name                 string `form:"name"`
    Password             string `form:"password"`
    PasswordConfirmation string `form:"password_confirmation"`
}

type UpdateUserRequest struct {
    ID                  int `form:"id"`
    Email               string `form:"email"`
    Name                string `form:"email"`
    Password            string `form:"password"`
}

type GenericUserResponse struct {
    ID    int       `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
}
