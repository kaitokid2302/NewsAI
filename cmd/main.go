package main

import (
	_ "github.com/kaitokid2302/NewsAI/docs"
	"github.com/kaitokid2302/NewsAI/internal/app"
	"github.com/kaitokid2302/NewsAI/internal/config"
)

func main() {
	config.InitAll()

	app.Run()

	// db := database.InitDatabase()
	// redisClient := redis.InitRedis()
	// userRepo := repository.NewUserRepo(db)
	// userService := userservice.NewUserService(userRepo, redisClient)

	// code, er := userService.SendEmail("truonglamthientai321@gmail.com")
	// fmt.Printf("er: %v\n", er)
	// fmt.Printf("code: %v\n", code)
}
