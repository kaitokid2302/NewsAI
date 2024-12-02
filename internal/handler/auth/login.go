package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/pkg/reponse"
	"github.com/kaitokid2302/NewsAI/pkg/request"
)

func (a *AuthHandler) Login(c *gin.Context) {
	var input request.LoginRequest

	if er := c.ShouldBind(&input); er != nil {
		reponse.ReponseOutput(c, reponse.LoginFail, "", nil)
		return
	}
	user, er := a.authService.Login(input.Email, input.Password)
	if er != nil {
		reponse.ReponseOutput(c, reponse.LoginFail, "", nil)
		return
	}
	token := a.jwtService.CreateToken(user.Email)
	user.Password = ""
	reponse.ReponseOutput(c, reponse.LoginSucess, "", gin.H{"user": user, "token": token})
}
