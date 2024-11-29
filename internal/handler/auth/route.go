package auth

import "github.com/gin-gonic/gin"

func (auth *AuthHandler) InitRoute(r *gin.RouterGroup) {
	r.POST("/register", auth.Register)
	r.POST("/verify", auth.VerifyOTP)
	r.POST("/verify/resend", auth.ResendOTP)
}
