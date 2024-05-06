package repository

import (
	"QueueOptimization/consts"
	"QueueOptimization/dtos"
	"QueueOptimization/models"
	"fmt"
	"log"
	"time"
)

func (r *repository) CreateQueue(queue models.Queue) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (place_id, age, predicted_time, iin) 
				VALUES ($1, $2, $3, $4) RETURNING id`, consts.QueueTable)
	row := r.db.QueryRow(query, queue.PlaceId, queue.Age, queue.PredictedTime, queue.IIN)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	query = fmt.Sprintf(`update %s set pending_time = pending_time + $1 where number = $2`, consts.PlaceTable)
	_, err := r.db.Query(query, queue.PredictedTime, queue.PlaceId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (r *repository) GetQueueInfo(iin string) (dtos.GetQueueInfoResponse, error) {
	var result dtos.GetQueueInfoResponse
	var created_time time.Time
	query := fmt.Sprintf("select created_time from %s where iin = $1", consts.QueueTable)
	err := r.db.Get(&created_time, query, iin)

	if err != nil {
		log.Print(err.Error())
		return dtos.GetQueueInfoResponse{}, err
	}

	log.Print(created_time)
	query = fmt.Sprintf(`select count(*) as people_count, sum(predicted_time) as time_count from %s where $1 > created_time`, consts.QueueTable)
	err = r.db.Get(&result, query, created_time)

	if err != nil {
		log.Print(err.Error())
	}
	return result, err
}
