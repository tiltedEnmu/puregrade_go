package httphandler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (h *HTTPHandler) AuthMiddleware(c *gin.Context) {
	header := strings.Split(c.GetHeader("Authorization"), " ")
	if (len(header) != 2) || (header[0] != "Bearer") {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	fmt.Print(header[1])

	id, err := h.services.Authorization.ParseAccessToken(header[1])
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "Token parse error: "+err.Error())
		return
	}

	c.Set("UserId", id)
}

func getUserId(c *gin.Context) (int64, error) {
	id, ok := c.Get("UserId")
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0, errors.New("UserId not found")
	}

	intId, ok := id.(int64)
	if !ok {
		c.AbortWithStatus(http.StatusUnauthorized)
		return 0, errors.New("UserId not found")
	}
	return intId, nil
}

func (h *HTTPHandler) CheckRoleMiddleware(c *gin.Context) {

}
