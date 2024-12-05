package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/aws"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	user2 "github.com/kaitokid2302/NewsAI/internal/handler/user"
	"github.com/kaitokid2302/NewsAI/internal/middleware"
	"github.com/kaitokid2302/NewsAI/internal/redis"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
	"github.com/kaitokid2302/NewsAI/internal/service/s3"
	"github.com/kaitokid2302/NewsAI/internal/service/user"

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

	session := aws.AwsInit()
	s3Service := s3.NewUploadFileS3Service(session)
	userService := user.NewUserService(s3Service, userRepo)
	userHandler := user2.NewUserHandler(userService)
	authHandler.InitRoute(r.Group("/auth"))
	userGroup := r.Group("/user")
	userGroup.Use(middleware.NewAuth(jwt.NewJWTService()).JWTverify())
	userHandler.InitRoute(userGroup)

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
