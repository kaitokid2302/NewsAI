package auth

import "github.com/gin-gonic/gin"

func (auth *AuthHandler) InitRoute(r *gin.RouterGroup) {
	r.POST("/register", auth.Register)
}
