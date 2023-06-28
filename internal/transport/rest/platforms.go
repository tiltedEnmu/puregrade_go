package httphandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDeletePlatformsDTO struct {
	ProductId   int64   `json:"productId"`
	PlatformsId []int64 `json:"platforms"`
}

func (h *HTTPHandler) AddPlatforms(c *gin.Context) {
	var input CreateDeletePlatformsDTO
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bind error: "+err.Error())
		return
	}

	if err := h.services.Product.AddPlatforms(input.ProductId, input.PlatformsId); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *HTTPHandler) DeletePlatforms(c *gin.Context) {
	var input CreateDeletePlatformsDTO
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bind error: "+err.Error())
		return
	}

	if err := h.services.Product.DeletePlatforms(input.ProductId, input.PlatformsId); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
