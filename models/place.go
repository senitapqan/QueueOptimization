package models

type Place struct {
	Id         int `json:"id"`
	Number     int `json:"place_name"`
	CategoryId int `json:"category_id"`
}
