package entity

type User struct {
    ID       int
	Email    string
	Name     string
	Password string
}

func (u *User) May(what string) bool {
    return false
}
