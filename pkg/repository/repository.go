package repository

import (
	"QueueOptimization/models"
	"QueueOptimization/dtos"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	CreateQueue(queue models.Queue) (int, error) 
	GetQueueInfo(iin string) (dtos.GetQueueInfoResponse, error)

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