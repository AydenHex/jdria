package api

type SendRequest struct {
	GameId  string `json:"gameId"`
	Message string `json:"message"`
}
