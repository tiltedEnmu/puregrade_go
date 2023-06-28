package httphandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) GetAllProducts(c *gin.Context) {
	var data puregrade.ProductFilter
	if err := c.BindQuery(&data); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	products, err := h.services.Product.GetAll(data)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, products)
}

func (h *HTTPHandler) GetOneProduct(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64) // string to int64
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

func (h *HTTPHandler) CreateProduct(c *gin.Context) {
	var input puregrade.CreateProductDTO
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Product.Create(input)
	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *HTTPHandler) DeleteProduct(c *gin.Context) {
	var id int64
	if err := c.BindJSON(&id); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Product.Delete(id); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(200)
}
