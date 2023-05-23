package services

import (
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/vdamery/jdria/internal/pkg/models"
	"github.com/vdamery/jdria/pkg/api"
	"github.com/vdamery/jdria/resources"
)

type InternalBootService interface {
	Test(name string) (*api.TestResponse, error)
}

// Enforces implementation of interface at compile time
var _ InternalBootService = (*InternalBootServiceImpl)(nil)

type InternalBootServiceImpl struct{}

func NewInternalBootService() *InternalBootServiceImpl {
	return &InternalBootServiceImpl{}
}

func (s *InternalBootServiceImpl) Test(name string) (*api.TestResponse, error) {
	var data models.TestData
	err := json.Unmarshal(resources.JsonExample, &data)
	if err != nil {
		return nil, errors.Wrap(err, "Test json.Unmarshal")
	}

	return &api.TestResponse{
		Message: "Hello " + name + " !",
		Data:    api.TestDataResponse(data),
	}, nil
}
