package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
)

type Hanlder struct {
	app         *gin.Engine
	authHandler *auth.AuthHandler
}

func NewHandler() *Hanlder {
	app := gin.Default()
	return &Hanlder{
		app:         app,
		authHandler: auth.NewAuthHandler(app),
	}
}
