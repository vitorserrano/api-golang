package middleware

import (
	"api/golang/internal/commons"
	"strings"

	"github.com/gin-gonic/gin"
)

const authorizationType = "Bearer"

func TokenValidate() gin.HandlerFunc {
	return func(c *gin.Context) {

		headerToken := c.GetHeader("Authorization")

		if len(headerToken) == 0 || !strings.Contains(headerToken, authorizationType) {
			c.AbortWithStatus(401)
			return
		}

		token := strings.Split(headerToken, authorizationType)[1]

		if err := commons.ValidateToken(token); err != nil {
			c.AbortWithStatus(401)
			return
		}
	}
}
