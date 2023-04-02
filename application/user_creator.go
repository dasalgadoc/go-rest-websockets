package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
)

type UserCreator struct {
	userRepository domain.UserRepository
}

func NewUserCreator(r domain.UserRepository) UserCreator {
	return UserCreator{
		userRepository: r,
	}
}

func (uc *UserCreator) Create(ctx context.Context, id, email, password string) error {
	var user = domain.NewUserFromPrimitives(id, email, password)
	return uc.userRepository.Insert(ctx, user)
}
