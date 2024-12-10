package user

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	reponse2 "github.com/kaitokid2302/NewsAI/internal/reponse"
	"github.com/kaitokid2302/NewsAI/internal/request"
	"github.com/kaitokid2302/NewsAI/internal/service/user"
)

type UserHandler struct {
	userService user.UserService
}

func NewUserHandler(userService user.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (uHandler *UserHandler) UpdateUser(c *gin.Context) {
	var input request.UpdateUserRequest
	if er := c.ShouldBind(&input); er != nil {
		reponse2.ReponseOutput(c, reponse2.UpdateUserFail, "", nil)
		return
	}
	var er error
	var u *database.User
	if input.Avatar != nil {
		file, _ := input.Avatar.Open()
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				reponse2.ReponseOutput(c, reponse2.UpdateUserFail, "", nil)
			}
		}(file)
		u, er = uHandler.userService.UpdateUser(c, input.Name, input.Avatar.Filename, &file)
	} else {
		u, er = uHandler.userService.UpdateUser(c, input.Name, input.Avatar.Filename, nil)
	}
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.UpdateUserFail, "", u)
		return
	}
	u.Password = ""
	reponse2.ReponseOutput(c, reponse2.UpdateUserSuccess, "", u)
}

func (uHandler *UserHandler) UserInfo(c *gin.Context) {
	email := c.GetString("email")
	user, er := uHandler.userService.GetUserInfo(email)
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.GetUserFail, "", nil)
		return
	}
	reponse2.ReponseOutput(c, reponse2.GetUserSuccess, "", user)
}
