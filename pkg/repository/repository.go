package repository

import (
	"QueueOptimization/models"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateQueue(queue models.Queue) (int, error) 
	/*GetAllQueues() ([]models.Queue, error) 
	GetQueueById(id int) (models.Queue, error)
	UpdateQueue(queue models.Queue) error
	DeleteQueue(id int) error */

	GetMostFreePlaceByCategoryId(id int) (int, error)

	PredictTime(age, id int) (int, error)
}

type repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &repository{
		db: db,
	}
}