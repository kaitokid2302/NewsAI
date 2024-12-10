package user

import (
	"github.com/gin-gonic/gin"
)

func (user *UserHandler) InitRoute(r *gin.RouterGroup) {
	// GET /user
	r.PUT("/update", user.UpdateUser)
	r.GET("/info", user.UserInfo)
}
