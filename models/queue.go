package models

type Queue struct {
	Id            int    `json:"id"`
	PlaceId       int    `json:"place_id"`
	Age           int    `json:"age"`
	PredictedTime int `json:"predicted_time"`
}

func NewQueue(time, age, id int) Queue {
	return Queue{
		PlaceId:       id,
		Age:           age,
		PredictedTime: time,
	}
}
