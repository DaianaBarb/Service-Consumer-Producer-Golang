package service

import (
	"project-golang/internal/domain/model"
)

type Processor interface {
	Process(updated *model.Model) error
}

type service struct {
}

func (s *service) Process(updated *model.Model) error {
	return nil

}

func NewService() Processor  {
	return &service{}
}


