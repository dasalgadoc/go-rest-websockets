package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
	"time"
)

type PostCreator struct {
	postRepository domain.PostRepository
}

func NewPostCreator(r domain.PostRepository) PostCreator {
	return PostCreator{
		postRepository: r,
	}
}

func (pc *PostCreator) Invoke(ctx context.Context, id, postContent, userId string, createdAt time.Time) error {
	var post = domain.NewPostFromPrimitives(id, postContent, userId, createdAt)
	return pc.postRepository.Insert(ctx, post)
}
