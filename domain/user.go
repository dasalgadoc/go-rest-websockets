package domain

type UserId string
type UserEmail string
type UserPassword string

type User struct {
	Id       UserId       `json:"id"`
	Email    UserEmail    `json:"email"`
	Password UserPassword `json:"password"`
}

func NewUserFromPrimitives(id, email, password string) User {
	return User{
		Id:       UserId(id),
		Email:    UserEmail(email),
		Password: UserPassword(password),
	}
}

func (u *User) IsUserEmpty() bool {
	return u.Email == "" || u.Password == ""
}
