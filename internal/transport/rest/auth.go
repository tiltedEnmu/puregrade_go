package httphandler

import (
	"log"
	"net/http"

	"github.com/ZaiPeeKann/puregrade"
	"github.com/gin-gonic/gin"
)

type singInDTO struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (h *HTTPHandler) singUp(c *gin.Context) {
	var input puregrade.User
	if err := c.BindJSON(&input); err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		log.Print(err.Error())
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *HTTPHandler) singIn(c *gin.Context) {
	var input singInDTO
	if err := c.BindJSON(&input); err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Username, input.Password)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
