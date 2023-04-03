package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
	"time"
)

type PostUpdater struct {
	postRepository domain.PostRepository
}

func NewPostUpdater(r domain.PostRepository) PostUpdater {
	return PostUpdater{
		postRepository: r,
	}
}

func (pu *PostUpdater) Invoke(ctx context.Context, id, postContent, userId string) error {
	var post = domain.NewPostFromPrimitives(id, postContent, userId, time.Now())
	return pu.postRepository.Update(ctx, post)
}
