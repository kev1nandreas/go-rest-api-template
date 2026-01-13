package main

import (
	"context"
	"log"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/kev1nandreas/go-rest-api-template/env"
	"github.com/kev1nandreas/go-rest-api-template/pkg/api"
	"github.com/kev1nandreas/go-rest-api-template/pkg/cache"
	"github.com/kev1nandreas/go-rest-api-template/pkg/database"
	"go.mongodb.org/mongo-driver/mongo"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8001
// @BasePath  /api/v1

// @securityDefinitions.apikey JwtAuth
// @in header
// @name Authorization

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-Key

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or error loading .env file")
	}

	appPort := env.GetEnvInt("APP_PORT", 8080)
	isLogging := env.GetEnvBool("APP_MONGO_LOGGING", false)
	isDebugging := env.GetEnvBool("APP_DEBUG", false)

	if isDebugging {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	redisClient := cache.NewRedisClient()
	db := database.NewDatabase()
	dbWrapper := &database.GormDatabase{DB: db}
	var mongo *mongo.Collection

	if isLogging {
		mongo = database.SetupMongoDB()
	}

	ctx := context.Background()
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := api.NewRouter(logger, mongo, dbWrapper, redisClient, &ctx)

	if err := r.Run(":" + strconv.Itoa(appPort)); err != nil {
		log.Fatal(err)
	}
}
