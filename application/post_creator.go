package application

import (
	"context"
	domain2 "dasalgadoc.com/rest-websockets/api/domain"
	"dasalgadoc.com/rest-websockets/domain"
	"time"
)

type PostCreator struct {
	postRepository domain.PostRepository
	websocketHub   *domain2.WebsocketHub
}

func NewPostCreator(r domain.PostRepository, hub *domain2.WebsocketHub) PostCreator {
	return PostCreator{
		postRepository: r,
		websocketHub:   hub,
	}
}

func (pc *PostCreator) Invoke(ctx context.Context, id, postContent, userId string, createdAt time.Time) error {
	var post = domain.NewPostFromPrimitives(id, postContent, userId, createdAt)
	err := pc.postRepository.Insert(ctx, post)
	if err != nil {
		return err
	}
	var postMessage = domain.WebsocketMessage{
		Type:    "post_created",
		Payload: post,
	}

	pc.websocketHub.Broadcast(postMessage, nil)
	return nil
}
