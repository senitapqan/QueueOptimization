package models

type Archive struct {
	Id int 
	Age int 
	CategoryId int
	PendingTime int
}

type Predictor struct {
	CategoryId  int    `json:"category_id"`
	AgeIn       int    `json:"age_in"`
	AgeOut      int    `json:"age_out"`
	AverageTime int `json:"average_time"`
}
