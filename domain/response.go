package domain

type Response struct {
	Message string `json:"message"`
	Status  bool   `json:"status"`
}
