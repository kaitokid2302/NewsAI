package auth

import "github.com/gin-gonic/gin"

type AuthHandler struct {
	App *gin.Engine
}

func NewAuthHandler(app *gin.Engine) *AuthHandler {
	return &AuthHandler{App: app}
}

func (authHandler *AuthHandler) InitRoute() {

}
