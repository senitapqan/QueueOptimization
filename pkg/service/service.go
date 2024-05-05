package service

import (
	//"QueueOptimization/models"
	"QueueOptimization/dtos"
	"QueueOptimization/pkg/repository"
)

type Service interface {
	CreateQueue(queue dtos.CreateQueueRequest) (int, error) 
	/*GetAllQueues() ([]models.Queue, error) 
	GetQueueById(id int) (models.Queue, error)
	UpdateQueue(queue models.Queue) error
	DeleteQueue(id int) error */
	
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