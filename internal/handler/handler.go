package handler

import (
	"Advertisement/internal/handler/metrics"
	"Advertisement/internal/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	metrics.RegisterStandardMetricsHandler()
	api := router.Group("/api")
	{
		ads := api.Group("/ads")
		{
			ads.POST("/", h.create)
			ads.GET("/", h.getAll)
			ads.GET("/:id", h.getById)
			ads.DELETE("/:id", h.delete)
			ads.PUT("/:id", h.update)
		}
	}
	return router
}
