package handler

import (
	"QueueOptimization/dtos"
	"log"
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

func (h *Handler) GetQueueInfo(c *gin.Context) {
	var input dtos.GetQueueInfoRequest
	err := ValidateIIN(c, &input)

	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	log.Print(input.IIN)

	info, err := h.service.GetQueueInfo(input.IIN)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, info)
}
