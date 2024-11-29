package main

import (
	"fmt"

	_ "github.com/kaitokid2302/NewsAI/docs"
	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/redis"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	userservice "github.com/kaitokid2302/NewsAI/internal/service/user"
)

func main() {
	config.InitAll()

	// app.Run()

	db := database.InitDatabase()
	redisClient := redis.InitRedis()
	userRepo := repository.NewUserRepo(db)
	userService := userservice.NewUserService(userRepo, redisClient)

	code, er := userService.SendEmail("truonglamthientai321@gmail.com")
	fmt.Printf("er: %v\n", er)
	fmt.Printf("code: %v\n", code)
}
