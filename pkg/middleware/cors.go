package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kev1nandreas/go-rest-api-template/env"
)

func Cors() gin.HandlerFunc {
	app_url := env.GetEnvString("APP_URL_PROD", "http://localhost:8080")
	allowedOrigins := []string{
		"http://127.0.0.1",
		"http://localhost",
		"http://localhost:3000",
		"http://localhost:8080",
		"http://localhost:5173",
		app_url,
	}

	return cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "https://github.com"
		//},
		MaxAge: 12 * time.Hour,
	})
}
