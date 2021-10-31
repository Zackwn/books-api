package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/zackwn/books-api/services"
)

func NewRequiredAuth(jwtService services.JWT) func(c *gin.Context) {
	return func(c *gin.Context) {
		rawToken := c.GetHeader("authorization")
		if rawToken == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		arrToken := strings.Split(rawToken, "Bearer ")
		if len(arrToken) != 2 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		token := arrToken[1]
		id, err := jwtService.Verify(token)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		c.Set("id", id)
		c.Next()
	}
}
