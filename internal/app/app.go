package app

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	"github.com/kaitokid2302/NewsAI/internal/redis"
	"github.com/kaitokid2302/NewsAI/internal/service"

	swaggerfiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run() {
	r := gin.Default()
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, url))

	redisClient := redis.InitRedis()
	emailService := service.NewEmailService(redisClient)
	authHandler := auth.NewAuthHandler(emailService, redisClient)

	authHandler.InitRoute(r.Group("/"))

	r.Run(":8080")
}
