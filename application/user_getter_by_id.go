package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
)

type UserGetterById struct {
	userRepository domain.UserRepository
}

func NewUserGetter(r domain.UserRepository) UserGetterById {
	return UserGetterById{
		userRepository: r,
	}
}

func (ug *UserGetterById) Invoke(ctx context.Context, id string) (domain.User, error) {
	return ug.userRepository.GetUserById(ctx, domain.UserId(id))
}
