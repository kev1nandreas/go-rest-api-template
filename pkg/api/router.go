package api

import (
	"context"
	"time"

	"github.com/kev1nandreas/go-rest-api-template/env"
	"github.com/kev1nandreas/go-rest-api-template/pkg/cache"
	"github.com/kev1nandreas/go-rest-api-template/pkg/database"
	"github.com/kev1nandreas/go-rest-api-template/pkg/middleware"

	docs "github.com/kev1nandreas/go-rest-api-template/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"

	"golang.org/x/time/rate"
)

func ContextMiddleware(bookRepository BookRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("appCtx", bookRepository)
		c.Next()
	}
}

func NewRouter(logger *zap.Logger, mongoCollection *mongo.Collection, db database.Database, redisClient cache.Cache, ctx *context.Context) *gin.Engine {
	isLogging := env.GetEnvBool("APP_MONGO_LOGGING", false)
	bookRepository := NewBookRepository(db, redisClient, ctx)
	userRepository := NewUserRepository(db, ctx)

	r := gin.Default()
	r.Use(ContextMiddleware(bookRepository))

	if isLogging {
		logger.Info("MongoDB logging is enabled")
		r.Use(middleware.Logger(logger, mongoCollection))
	} else {
		logger.Info("MongoDB logging is disabled")
		r.Use(gin.Logger())
	}

	if gin.Mode() == gin.ReleaseMode {
		r.Use(middleware.Security())
		r.Use(middleware.Xss())
	}
	r.Use(middleware.Cors())
	r.Use(middleware.RateLimiter(rate.Every(1*time.Minute), 60)) // 60 requests per minute

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/", bookRepository.Healthcheck)
		v1.GET("/books", middleware.APIKeyAuth(), bookRepository.FindBooks)
		v1.POST("/books", middleware.APIKeyAuth(), middleware.JWTAuth(), bookRepository.CreateBook)
		v1.GET("/books/:id", middleware.APIKeyAuth(), bookRepository.FindBook)
		v1.PUT("/books/:id", middleware.APIKeyAuth(), bookRepository.UpdateBook)
		v1.DELETE("/books/:id", middleware.APIKeyAuth(), bookRepository.DeleteBook)

		v1.POST("/login", middleware.APIKeyAuth(), userRepository.LoginHandler)
		v1.POST("/register", middleware.APIKeyAuth(), userRepository.RegisterHandler)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
