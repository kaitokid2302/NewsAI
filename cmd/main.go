package main

import (
	"fmt"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/service"
)

func main() {
	config.InitAll()
	database.InitDatabase()

	email := service.NewEmailService()
	er := email.SendEmail("dinhtruonglam2001ctn@gmail.com")
	fmt.Printf("er: %v\n", er)
}
