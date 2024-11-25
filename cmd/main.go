package main

import (
	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/database"
)

func main() {
	config.InitAll()
	database.InitDatabase()
}
