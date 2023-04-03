package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
)

type PostDeleter struct {
	postRepository domain.PostRepository
}

func NewPostDeleter(r domain.PostRepository) PostDeleter {
	return PostDeleter{
		postRepository: r,
	}
}

func (pd *PostDeleter) Invoke(ctx context.Context, postId, userId string) error {
	return pd.postRepository.Delete(ctx, domain.PostId(postId), domain.UserId(userId))
}
