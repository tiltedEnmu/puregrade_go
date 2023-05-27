package httphandler

import (
	"github.com/ZaiPeeKann/puregrade/internal/service"
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
		review.DELETE("/", h.AuthMiddleware, h.DeleteReview)
	}
	product := router.Group("/product")
	{
		product.POST("/", h.AuthMiddleware, h.CreateProduct)
		product.GET("/", h.GetAllProducts)
		product.GET("/:id", h.GetOneProduct)
		product.DELETE("/", h.DeleteProduct)
		genres := product.Group("/genres")
		{
			genres.POST("/", h.AuthMiddleware, h.AddGenres)
			genres.DELETE("/", h.AuthMiddleware, h.DeleteGenres)
		}
		platforms := product.Group("/platforms")
		{
			platforms.POST("/", h.AuthMiddleware, h.AddPlatforms)
			platforms.DELETE("/", h.AuthMiddleware, h.DeletePlatforms)
		}
	}
	user := router.Group("/user")
	{
		user.GET("/:id", h.AuthMiddleware, h.GetProfile)
		user.DELETE("/:id", h.AuthMiddleware, h.DeleteUser)
		followers := user.Group("/followers")
		{
			followers.POST("/:id", h.AuthMiddleware, h.FollowUser)
			followers.DELETE("/:id", h.AuthMiddleware, h.UnfollowUser)
		}
	}

	return router
}
