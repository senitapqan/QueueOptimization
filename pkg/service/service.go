package service

import (
	//"QueueOptimization/models"
	"QueueOptimization/dtos"
	"QueueOptimization/pkg/repository"
)

type Service interface {
	CreateQueue(queue dtos.CreateQueueRequest) (int, error) 
	
	GetMostFreePlaceByCategoryId(id int) (int, error)
}

type service struct {
	repos repository.Repository
}

func NewService(r repository.Repository) Service {
	return &service{
		repos: r,
	}
}