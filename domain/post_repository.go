package domain

import "context"

var PostRepositoryImplementation PostRepository

type PostRepository interface {
	Insert(ctx context.Context, post Post) error
	GetPostById(ctx context.Context, postId PostId) (Post, error)
	Update(ctx context.Context, post Post) error
	Delete(ctx context.Context, postId PostId, userId UserId) error
	Close() error
}

func SetPostRepository(repository PostRepository) {
	PostRepositoryImplementation = repository
}
