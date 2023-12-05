package model

type CreateUserRequest struct {
    Email                string `form:"email"`
    Name                 string `form:"name"`
    Password             string `form:"password"`
    PasswordConfirmation string `form:"password_confirmation"`
}

type GenericUserResponse struct {
    ID    int       `json:"id"`
    Name  string    `json:"name"`
    Email string    `json:"email"`
}
