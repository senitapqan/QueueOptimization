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

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, id)
}

// func (h *Handler) GetQueueInfo(c *gin.Context) {
// 	var input dtos.GetQueueInfoRequest
// 	if err := c.BindJSON(&input); err != nil {

func (h *Handler) GetAllQueues(c *gin.Context) {
	queues, err := h.service.getAllQueues()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, queues)
}

func (h *Handler) GetQueueById(c *gin.Context) {
	_, err := ValidateId(c)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

}
