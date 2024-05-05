package dtos

type CreateQueueRequest struct {
	IIN        string `json:"iin"`
	CategoryId int    `json:"category_id"`
}

type GetQueueInfoRequest struct {
	IIN string `json:"iin"`
}

type GetQueueInfoResponse struct {
	PeopleCount int `json:"people_count"`
	TimeCount   int `json:"time_count"`
}
