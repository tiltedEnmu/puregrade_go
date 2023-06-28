package httphandler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateDeleteGenresDTO struct {
	ProductId int64   `json:"productId"`
	GenresId  []int64 `json:"genres"`
}

func (h *HTTPHandler) AddGenres(c *gin.Context) {
	var input CreateDeleteGenresDTO
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bind error: "+err.Error())
		return
	}

	if err := h.services.Product.AddGenres(input.ProductId, input.GenresId); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *HTTPHandler) DeleteGenres(c *gin.Context) {
	var input CreateDeleteGenresDTO
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, "Bind error: "+err.Error())
		return
	}

	if err := h.services.Product.DeleteGenres(input.ProductId, input.GenresId); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
