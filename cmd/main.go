package main

import (
	_ "github.com/kaitokid2302/NewsAI/docs"
	"github.com/kaitokid2302/NewsAI/internal/app"
	"github.com/kaitokid2302/NewsAI/internal/config"
)

func main() {
	config.InitAll()

	app.Run()
}
