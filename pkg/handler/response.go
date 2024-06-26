package handler

import (
	"QueueOptimization/dtos"
	"errors"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func newErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func ValidateId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return -1, err
	}

	if id <= 0 {
		return -1, errors.New("id cannot be negative")
	}
	return id, nil
}

func ValidateIIN(c *gin.Context, input *dtos.GetQueueInfoRequest) error {
	iin := c.Query("iin")

	log.Print("g" + iin)
	input.IIN = iin

	return nil
}
