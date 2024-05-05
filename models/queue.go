package models

type Queue struct {
	Id      int    `json:"id"`
	PlaceId int    `json:"place_id"`
	Age     int    `json:"age"`
	Time    string `json:"time"`
}

func NewQueue(time string, age, id int) Queue {
	return Queue{
		PlaceId: id,
		Age: age,
		Time: time,
	}
}
