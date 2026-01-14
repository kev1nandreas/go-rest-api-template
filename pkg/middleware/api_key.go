package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kev1nandreas/go-rest-api-template/env"
	"github.com/kev1nandreas/go-rest-api-template/pkg/auth"
	"github.com/kev1nandreas/go-rest-api-template/pkg/response"
)

func APIKeyAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		if apiKey == env.GetEnvString("API_SECRET_KEY", auth.GenerateRandomKey()) {
			c.Next()
		} else {
			response.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Invalid API Key",
			).Send(c)
			c.Abort()
		}
	}
}
