package httphandler

import (
	"github.com/ZaiPeeKann/auth-service_pg/internal/service"
	"github.com/gin-gonic/gin"
)

type HTTPHandler struct {
	services *service.Service
}

func NewHTTPHandler(services *service.Service) *HTTPHandler {
	return &HTTPHandler{services: services}
}

func (h *HTTPHandler) InitRoutes() *gin.Engine {
	router := gin.New()
	auth := router.Group("/auth")
	{
		auth.POST("/sing-in", h.singIn)
		auth.POST("/sing-up", h.singUp)
	}
	review := router.Group("/review")
	{
		review.POST("/", h.AuthMiddleware, h.CreateReview)
		review.GET("/", h.GetAllReviews)
		review.GET("/:id", h.GetOneReview)
		review.PATCH("/:id", h.AuthMiddleware, h.UpdateReview)
		review.DELETE("/:id", h.AuthMiddleware, h.DeleteReview)
	}

	return router
}
