package domain

import "time"

type PostId string
type PostContent string
type PostDateCreation time.Time

type Post struct {
	Id          PostId           `json:"id"`
	PostContent PostContent      `json:"post_content"`
	CreatedAt   PostDateCreation `json:"created_at"`
	UserId      UserId           `json:"user_id"`
}

func NewPostFromPrimitives(id, postContent, userId string, createdAt time.Time) Post {
	return Post{
		Id:          PostId(id),
		PostContent: PostContent(postContent),
		CreatedAt:   PostDateCreation(createdAt),
		UserId:      UserId(userId),
	}
}
