package middleware

import (
	"net/http"
	"strings"

	"github.com/kev1nandreas/go-rest-api-template/pkg/auth"
	"github.com/kev1nandreas/go-rest-api-template/pkg/response"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			response.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Authorization header missing",
			).Send(c)
			c.Abort()
			return
		}

		if !strings.HasPrefix(header, BearerSchema) {
			response.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Invalid authorization header format",
			).Send(c)
			c.Abort()
			return
		}

		tokenStr := header[len(BearerSchema):]
		claims := &auth.Claims{}

		token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
			return auth.JwtKey, nil
		})

		if err != nil {
			response.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Invalid token: "+err.Error(),
			).Send(c)
			c.Abort()
			return
		}

		if !token.Valid {
			response.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthorized",
				"Token is not valid",
			).Send(c)
			c.Abort()
			return
		}

		c.Set("username", claims.Username)
		c.Next()
	}
}
