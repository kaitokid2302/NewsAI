package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	"github.com/kaitokid2302/NewsAI/internal/redis"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	r := gin.Default()
	// cors all
	r.Use(cors.Default())
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	redisClient := redis.InitRedis()
	db := database.InitDatabase()
	jwtService := jwt.NewJWTService()
	userRepo := repository.NewUserRepo(db)
	authService := authService.NewAuthService(userRepo, redisClient)

	authHandler := auth.NewAuthHandler(authService, jwtService)

	authHandler.InitRoute(r.Group("/auth"))

	r.Run(":8080")
}
