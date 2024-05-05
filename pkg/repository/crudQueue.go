package repository;

import (
	"QueueOptimization/models"
	"database/sql"
)

type QueueRepository struct {
	db *sql.DB
}

func NewQueueRepository(db *sql.DB) *QueueRepository {
	return &QueueRepository{db: db}
}

func (r *QueueRepository) CreateQueue(queue models.Queue) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	var id int
	row := tx.QueryRow("INSERT INTO queue (place_id, age, time) VALUES ($1, $2, $3) RETURNING id", queue.PlaceId, queue.Age, queue.Time)
	if err := row.Scan(&id); err != nil {
		tx.Rollback
		return 0, err
	}

	return id, tx.Commit()
}

func (r *QueueRepository) GetAllQueues() ([]models.Queue, error) {
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

func (r *QueueRepository) GetQueueById(id int) (models.Queue, error) {
	row := r.db.QueryRow("SELECT * FROM queue WHERE id = $1", id)

	queue := models.Queue{}
	err := row.Scan(&queue.Id, &queue.PlaceId, &queue.Age, &queue.Time)
	if err != nil {
		return models.Queue{}, err
	}

	return queue, nil
}

func (r *QueueRepository) UpdateQueue(queue models.Queue) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("UPDATE queue SET place_id = $1, age = $2, time = $3 WHERE id = $4", queue.PlaceId, queue.Age, queue.Time, queue.Id)
	if err != nil {
		tx.Rollback
		return err
	}

	return tx.Commit()
}

func (r *QueueRepository) DeleteQueue(id int) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("DELETE FROM queue WHERE id = $1", id)
	if err != nil {
		tx.Rollback
		return err
	}

	return tx.Commit()
}