package user

import (
	"github.com/gin-gonic/gin"
)

func (user *UserHandler) InitRoute(r *gin.RouterGroup) {
	// GET /user
	r.Use(user.jwtMiddleware.JWTverify())
	r.PUT("/user", user.UpdateUser)

}
