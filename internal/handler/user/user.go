package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/middleware"
	"github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/user"
)

type UserHandler struct {
	userService   user.UserService
	authService   auth.AuthService
	jwtMiddleware middleware.Auth
}

func NewUserHandler(userService user.UserService, authService auth.AuthService, jwtMiddleware middleware.Auth) *UserHandler {
	return &UserHandler{
		userService:   userService,
		authService:   authService,
		jwtMiddleware: jwtMiddleware,
	}
}

func (uHandler *UserHandler) UpdateUser(c *gin.Context) {
	// TODO: implement this
}
