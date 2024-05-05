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
	return router
}