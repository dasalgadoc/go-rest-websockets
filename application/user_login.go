package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
)

type UserLogin struct {
	userRepository domain.UserRepository
}

func NewUserLogin(r domain.UserRepository) UserLogin {
	return UserLogin{
		userRepository: r,
	}
}

func (ul *UserLogin) Do(ctx context.Context, email string) (domain.User, error) {
	return ul.userRepository.GetUserByEmail(ctx, domain.UserEmail(email))
}
