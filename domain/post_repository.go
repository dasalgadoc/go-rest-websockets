package domain

import "context"

var PostRepositoryImplementation PostRepository

type PostRepository interface {
	Insert(ctx context.Context, post Post) error
	Close() error
}

func SetPostRepository(repository PostRepository) {
	PostRepositoryImplementation = repository
}
