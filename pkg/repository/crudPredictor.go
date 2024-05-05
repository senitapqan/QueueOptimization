package repository

import (
	"QueueOptimization/consts"
	"fmt"
	"log"
)

func (r *repository) PredictTime(id, age int) (int, error) {
	var result int
	query := fmt.Sprintf(`select average_time from %s where age_in <= $1 and $2 <= age_out and category_id = $3`, consts.PredictorTable)
	log.Print(query)
	err := r.db.Get(&result, query, age, age, id)
	if err != nil {
		return 0, err
	}
	return result, nil
}
