package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
)

type Handler struct {
	app         *gin.Engine
	authHandler *auth.AuthHandler
}

func NewHandler() *Handler {
	app := gin.Default()
	return &Handler{
		app:         app,
		authHandler: auth.NewAuthHandler(app),
	}
}

func (handler *Handler) Run() {
	handler.authHandler.InitRoute()
}
