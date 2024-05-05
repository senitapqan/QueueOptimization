package dtos

type CreateQueueRequest struct {
	IIN        string `json:""`
	CategoryId int    `json:"category_id"`
}
