package repository

import (
	"QueueOptimization/consts"
	"QueueOptimization/models"
	"fmt"
)

func (r *repository) CreateQueue(queue models.Queue) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (place_id, age, predicted_time, iin) 
				VALUES ($1, $2, $3, $4) RETURNING id`, consts.QueueTable)
	row := r.db.QueryRow(query, queue.PlaceId, queue.Age, queue.PredictedTime, queue.IIN)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r * repository) GetAllQueues() ([]models.Queue, error) {
	var result []models.Queue
	query := fmt.Sprintf(`select id, iin from %s`, consts.QueueTable)
	err := r.db.Select(&result, query)
	
	return result, err
}