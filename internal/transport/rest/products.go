package httphandler

import (
	"net/http"
	"strconv"

	puregrade "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetAllProducts(c *gin.Context) {
	var queryParams puregrade.ProductFilter
	filter := make(map[string]string)
	if err := c.BindQuery(&queryParams); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if queryParams.Genre != "" {
		filter["genre"] = queryParams.Genre
	}
	if queryParams.Platform != "" {
		filter["platform"] = queryParams.Platform
	}

	products, err := h.services.Product.GetAll(queryParams.Page, filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *HTTPHandler) GetOneProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	product, err := h.services.Product.GetOneByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)
}
