package models

type Player struct {
	Id string `json:"id"`
}

type Game struct {
	Id      string   `json:"id"`
	Players []Player `json:"players"`
	History []string `json:"history"`
}
