package domain

import "context"

type UserRepository interface {
	Insert(ctx context.Context, user User) error
	GetUserById(ctx context.Context, id UserId) (User, error)
	GetUserByEmail(ctx context.Context, email UserEmail) (User, error)
	Close() error
}

var UserRepositoryImplementation UserRepository

func SetRepository(repository UserRepository) {
	UserRepositoryImplementation = repository
}

func Insert(ctx context.Context, user User) error {
	return UserRepositoryImplementation.Insert(ctx, user)
}

func GetUserById(ctx context.Context, id UserId) (User, error) {
	return UserRepositoryImplementation.GetUserById(ctx, id)
}

func GetUserByEmail(ctx context.Context, email UserEmail) (User, error) {
	return UserRepositoryImplementation.GetUserByEmail(ctx, email)
}

func Close() error {
	return UserRepositoryImplementation.Close()
}
