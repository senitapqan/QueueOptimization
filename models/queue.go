package models

type Queue struct {
	Id      int    `json:"id"`
	PlaceId int    `json:"place_id"`
	Age     string `json:"age"`
	Time    string `json:"time"`
}
