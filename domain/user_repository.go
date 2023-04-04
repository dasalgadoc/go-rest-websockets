package domain

import "context"

type UserRepository interface {
	Insert(ctx context.Context, user User) error
	GetUserById(ctx context.Context, id UserId) (User, error)
	GetUserByEmail(ctx context.Context, email UserEmail) (User, error)
	Close() error
}
