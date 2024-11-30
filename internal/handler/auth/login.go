package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *AuthHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required" form:"email"`
		Password string `json:"password" binding:"required" form:"password"`
	}
	var output Response

	if er := c.ShouldBind(&input); er != nil {

		output = Response{
			StatusCode: http.StatusBadRequest,
			Er:         er.Error(),
			Message:    "Username or password incorrect",
		}
		c.JSON(output.StatusCode, output)
		return
	}
	user, er := a.userService.Login(input.Email, input.Password)
	if er != nil {
		output = Response{
			StatusCode: http.StatusInternalServerError,
			Er:         er.Error(),
			Message:    "Username or password incorrect",
		}
		c.JSON(output.StatusCode, output)
		return
	}
	token := a.jwtService.CreateToken(user.Email)
	user.Password = ""
	output = Response{
		StatusCode: http.StatusOK,
		Message:    "Login success",
		JwtToken:   token,
		Data:       *user,
	}
	c.JSON(output.StatusCode, output)
}
