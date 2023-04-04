package domain

type Server interface {
	GetConfig() *Config
	GetHub() *WebsocketHub
}
