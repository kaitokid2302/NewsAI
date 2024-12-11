package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	user2 "github.com/kaitokid2302/NewsAI/internal/handler/user"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/aws"
	crobjob2 "github.com/kaitokid2302/NewsAI/internal/infrastructure/crobjob"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
	markdown2 "github.com/kaitokid2302/NewsAI/internal/infrastructure/markdown"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/redis"
	"github.com/kaitokid2302/NewsAI/internal/middleware"
	"github.com/kaitokid2302/NewsAI/internal/repository/article"
	user3 "github.com/kaitokid2302/NewsAI/internal/repository/user"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	crobjob3 "github.com/kaitokid2302/NewsAI/internal/service/crobjob"
	elastic2 "github.com/kaitokid2302/NewsAI/internal/service/elastic"
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
	_ = elastic.InitElasticSearch()

	redisClient := redis.InitRedis()
	db := database.InitDatabase()
	jwtService := jwt.NewJWTService()
	userRepo := user3.NewUserRepo(db)
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

	articleRepo := article.NewArticleRepo(db)
	markdown := markdown2.NewMarkdown()
	elasticClient := elastic.InitElasticSearch()
	elasticService := elastic2.NewElasticService(elasticClient)
	crobjobservice := crobjob3.NewCronJobArticleService(articleRepo, markdown, elasticService)
	crobjob := crobjob2.NewCrobjob(*crobjobservice)
	go func() {
		crobjob.Run()
	}()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
