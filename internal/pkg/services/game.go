package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/vdamery/jdria/internal/pkg/models"
	"github.com/vdamery/jdria/pkg/api"
)

type GameService interface {
	StartGame() (*models.Game, error)
	Send(request api.SendRequest) (*models.Game, error)
}

// Enforces implementation of interface at compile time
var _ GameService = (*GameServiceImpl)(nil)

type GameServiceImpl struct {
	Games map[string]models.Game
}

func NewInternalBootService() *GameServiceImpl {
	return &GameServiceImpl{
		Games: make(map[string]models.Game),
	}
}

func (s *GameServiceImpl) StartGame() (*models.Game, error) {
	history := make([]string, 0)
	gameId := uuid.New().String()
	game := models.Game{
		Id: gameId,
		Players: []models.Player{
			{
				Id: uuid.New().String(),
			},
		},
		History: history,
	}
	s.Games[gameId] = game

	return &game, nil
}

func (s *GameServiceImpl) Send(request api.SendRequest) (*models.Game, error) {
	game, ok := s.Games[request.GameId]
	if !ok {
		return nil, errors.New("game not found")
	}
	game.History = append(game.History, request.Message)
	s.Games[request.GameId] = game
	return &game, nil
}
