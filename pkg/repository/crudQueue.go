package repository

import (
	"QueueOptimization/models"
	"QueueOptimization/consts"
	"fmt"
)

func (r *repository) CreateQueue(queue models.Queue) (int, error) {
	var id int
	query := fmt.Sprintf(`INSERT INTO %s (place_id, age, predicted_time) 
				VALUES ($1, $2, $3) RETURNING id`, consts.QueueTable)
	row :=r.db.QueryRow(query, )
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

}

func (r *repository) GetAllQueues() ([]models.Queue, error) {
	rows, err := r.db.Query("SELECT * FROM queue")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	queues := make([]models.Queue, 0)
	for rows.Next() {
		queue := models.Queue{}
		err := rows.Scan(&queue.Id, &queue.PlaceId, &queue.Age, &queue.Time)
		if err != nil {
			return nil, err
		}
		queues = append(queues, queue)
	}

	return queues, nil
}

func (r *repository) GetQueueById(id int) (models.Queue, error) {
	row := r.db.QueryRow("SELECT * FROM queue WHERE id = $1", id)

	queue := models.Queue{}
	err := row.Scan(&queue.Id, &queue.PlaceId, &queue.Age, &queue.Time)
	if err != nil {
		return models.Queue{}, err
	}

	return queue, nil
}

func (r *repository) UpdateQueue(queue models.Queue) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE queue SET place_id = $1, age = $2, time = $3 WHERE id = $4", queue.PlaceId, queue.Age, queue.Time, queue.Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (r *repository) DeleteQueue(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM queue WHERE id = $1", id)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}