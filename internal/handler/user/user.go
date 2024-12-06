package user

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database/model"
	"github.com/kaitokid2302/NewsAI/internal/service/user"
	"github.com/kaitokid2302/NewsAI/pkg/reponse"
	"github.com/kaitokid2302/NewsAI/pkg/request"
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
		reponse.ReponseOutput(c, reponse.UpdateUserFail, "", nil)
		return
	}
	var er error
	var u *model.User
	if input.Avatar != nil {
		file, _ := input.Avatar.Open()
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				reponse.ReponseOutput(c, reponse.UpdateUserFail, "", nil)
			}
		}(file)
		u, er = uHandler.userService.UpdateUser(c, input.Name, input.Avatar.Filename, &file)
	} else {
		u, er = uHandler.userService.UpdateUser(c, input.Name, input.Avatar.Filename, nil)
	}
	if er != nil {
		reponse.ReponseOutput(c, reponse.UpdateUserFail, "", u)
		return
	}
	u.Password = ""
	reponse.ReponseOutput(c, reponse.UpdateUserSuccess, "", u)
}
