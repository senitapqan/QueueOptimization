package service

import (
	"QueueOptimization/models"
	"QueueOptimization/repository"
)

type CRUDQueueService struct {
	repo *repository.QueueRepository
}

func NewCRUDQueueService(repo *repository.QueueRepository) *CRUDQueueService {
	return &CRUDQueueService{repo: repo}
}

func (s *CRUDQueueService) CreateQueue(queue models.Queue) (int, error) {
	return s.repo.CreateQueue(queue)
}

func (s *CRUDQueueService) GetAllQueues() ([]models.Queue, error) {
	return s.repo.GetAllQueues()
}

func (s *CRUDQueueService) GetQueueById(id int) (models.Queue, error) {
	return s.repo.GetQueueById(id)
}

func (s *CRUDQueueService) UpdateQueue(queue models.Queue) error {
	return s.repo.UpdateQueue(queue)
}

func (s *CRUDQueueService) DeleteQueue(id int) error {
	return s.repo.DeleteQueue(id)
}
