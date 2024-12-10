package auth

import (
	"github.com/gin-gonic/gin"
	reponse2 "github.com/kaitokid2302/NewsAI/internal/reponse"
	"github.com/kaitokid2302/NewsAI/internal/request"
)

func (a *AuthHandler) Login(c *gin.Context) {
	var input request.LoginRequest

	if er := c.ShouldBind(&input); er != nil {
		reponse2.ReponseOutput(c, reponse2.LoginFail, er.Error(), nil)
		return
	}
	user, er := a.authService.Login(input.Email, input.Password)
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.LoginFail, er.Error(), nil)
		return
	}
	token := a.jwtService.CreateToken(user.Email)
	user.Password = ""
	reponse2.ReponseOutput(c, reponse2.LoginSucess, "", gin.H{"user": user, "token": token})
}
