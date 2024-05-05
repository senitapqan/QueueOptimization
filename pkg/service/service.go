package service

import (
	"QueueOptimization/pkg/repository"
)

type Service interface {
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}