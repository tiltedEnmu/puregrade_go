package httphandler

import (
	"net/http"
	"strconv"

	puregrade "github.com/ZaiPeeKann/auth-service_pg"
	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) CreateReview(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var input puregrade.Review
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	input.Author.Id = id

	reviewId, err := h.services.Review.Create(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": reviewId,
	})
}

func (h *HTTPHandler) GetAllReviews(c *gin.Context) {
	var queryParams puregrade.RewiewFilter
	if err := c.BindQuery(&queryParams); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	reviews, err := h.services.Review.GetAll(queryParams.Page, queryParams.ProductId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, reviews)
}

func (h *HTTPHandler) GetOneReview(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	review, err := h.services.Review.GetOneByID(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, review)
}

func (h *HTTPHandler) UpdateReview(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var input puregrade.Review
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if id != input.Author.Id {
		c.AbortWithStatusJSON(http.StatusForbidden, err.Error())
		return
	}

	if err := h.services.Review.Update(input.Id, input.Title, input.Body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}

func (h *HTTPHandler) DeleteReview(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	var reviewId int
	if err := c.BindJSON(&reviewId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.services.Review.Delete(reviewId, userId); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Status(http.StatusOK)
}
