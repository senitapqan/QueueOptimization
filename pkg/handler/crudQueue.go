package handler

import (
	"QueueOptimization/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)


func (h *Handler) CreateQueue(c *gin.Context) {
	var input dtos.CreateQueueRequest
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.CreateQueue(input)
}

func (h *Handler) GetAllQueues(c *gin.Context) {
	
}

func (h *Handler) GetQueueById(c *gin.Context) {
	_, err := ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

func (h *Handler) UpdateQueue(c *gin.Context) {
	
}

func (h *Handler) DeleteQueue(c *gin.Context) {
	id, err := ValidateId()
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
}

