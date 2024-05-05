package models;

type Place struct {
	Id int64 `json:"id"`
	PlaceName string `json:"place_name"`
	CategoryId int64 `json:"category_id"`
}