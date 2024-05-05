package handler

import (
	"github.com/gin-gonic/gin"
	"QueueOptimization/pkg/service"
)

type Handler struct {
	service service.Service
}

func NewHandler(serv service.Service) *Handler {
	return &Handler{
		service: serv,
	}
}

func (h Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	queue := router.Group("queue")
	{
		queue.POST("/", h.CreateQueue),
		queue.GET("/all", h.GetAllQueues),
	}

	return router
}