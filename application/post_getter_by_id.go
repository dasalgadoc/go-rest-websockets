package application

import (
	"context"
	"dasalgadoc.com/rest-websockets/domain"
)

type PostGetterById struct {
	postRepository domain.PostRepository
}

func NewPostGetter(r domain.PostRepository) PostGetterById {
	return PostGetterById{
		postRepository: r,
	}
}

func (pg *PostGetterById) Invoke(ctx context.Context, id string) (domain.Post, error) {
	return pg.postRepository.GetPostById(ctx, domain.PostId(id))
}
