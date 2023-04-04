package domain

import "dasalgadoc.com/rest-websockets/domain"

type Server interface {
	GetConfig() *Config
	GetHub() *WebsocketHub
	GetUserRepository() domain.UserRepository
	GetPostRepository() domain.PostRepository
}
