package main

import (
	"fmt"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/service/otp"
)

func main() {
	config.InitAll()
	database.InitDatabase()

	email := otp.NewEmail()
	er := email.SendEmail("dinhtruonglam2001ctn@gmail.com")
	fmt.Printf("er: %v\n", er)
}
