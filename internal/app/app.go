package app

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	article2 "github.com/kaitokid2302/NewsAI/internal/handler/article"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	topic3 "github.com/kaitokid2302/NewsAI/internal/handler/topic"
	user2 "github.com/kaitokid2302/NewsAI/internal/handler/user"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/aws"
	crobjob2 "github.com/kaitokid2302/NewsAI/internal/infrastructure/crobjob"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
	markdown2 "github.com/kaitokid2302/NewsAI/internal/infrastructure/markdown"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/redis"
	"github.com/kaitokid2302/NewsAI/internal/middleware"
	"github.com/kaitokid2302/NewsAI/internal/repository/article"
	topic2 "github.com/kaitokid2302/NewsAI/internal/repository/topic"
	user3 "github.com/kaitokid2302/NewsAI/internal/repository/user"
	article3 "github.com/kaitokid2302/NewsAI/internal/service/article"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	crobjob3 "github.com/kaitokid2302/NewsAI/internal/service/crobjob"
	elastic2 "github.com/kaitokid2302/NewsAI/internal/service/elastic"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
	"github.com/kaitokid2302/NewsAI/internal/service/s3"
	"github.com/kaitokid2302/NewsAI/internal/service/topic"
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
	topicRepo := topic2.NewTopicRepository(db)
	topicService := topic.NewTopicService(userService, topicRepo)
	topicHandler := topic3.NewTopicHandler(topicService)
	topicGroup := r.Group("/topic")
	topicGroup.Use(middleware.NewAuth(jwt.NewJWTService()).JWTverify())
	topicHandler.InitRoute(topicGroup)

	articleGroup := r.Group("/article")
	articleGroup.Use(middleware.NewAuth(jwt.NewJWTService()).JWTverify())
	articleService := article3.NewArticleService(articleRepo, userRepo, topicRepo)
	articleHandler := article2.NewArticleHandler(articleService)
	articleHandler.InitRoute(articleGroup)
	crobjob.Run()

	err := r.Run(":8080")
	if err != nil {
		return
	}
}
