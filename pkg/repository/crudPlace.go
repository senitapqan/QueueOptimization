package repository

import (
	"QueueOptimization/consts"
	"errors"
	"fmt"
)

func (r *repository) GetMostFreePlaceByCategoryId(category_id int) (int, error) {
	var id int
	query := fmt.Sprintf("Select id from %s where category_id = $1 order by pending_time limit 1 ", consts.PlaceTable)
	err := r.db.Get(&id, query, category_id)
	if err != nil {
		return 0, errors.New("something went wrong in request")
	}
	return id, nil
}