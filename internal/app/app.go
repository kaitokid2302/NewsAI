package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	"github.com/kaitokid2302/NewsAI/internal/redis"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	userservice "github.com/kaitokid2302/NewsAI/internal/service/user"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	r := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	redisClient := redis.InitRedis()
	db := database.InitDatabase()
	userRepo := repository.NewUserRepo(db)
	userService := userservice.NewUserService(userRepo, redisClient)
	authHandler := auth.NewAuthHandler(userService)

	authHandler.InitRoute(r.Group("/auth"))

	r.Run(":8080")
}
