package repository

import (
	"QueueOptimization/consts"
	"fmt"
)

func (r *repository) PredictTime(age, id int) (int, error) {
	var result int
	query := fmt.Sprintf(`Select from %s (average_time) where age_in <= $1 and $2 <= age_out`, consts.PredictorTable)
	err := r.db.Get(&result, query, age, age)
	if err != nil {
		return 0, err
	}
	return result, nil
}