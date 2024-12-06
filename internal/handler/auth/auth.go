package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database/model"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
	"github.com/kaitokid2302/NewsAI/pkg/reponse"
	"github.com/kaitokid2302/NewsAI/pkg/request"
	"github.com/peteprogrammer/go-automapper"
)

type AuthHandler struct {
	authService authService.AuthService
	jwtService  jwt.JWTservice
}

func NewAuthHandler(authService authService.AuthService, jwtService jwt.JWTservice) *AuthHandler {
	return &AuthHandler{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (auth *AuthHandler) Register(c *gin.Context) {
	var registerRequest request.RegisterRequest
	if er := c.ShouldBind(&registerRequest); er != nil {
		reponse.ReponseOutput(c, reponse.RegisterFail, "", nil)
		return
	}
	var user model.User
	automapper.MapLoose(registerRequest, &user)
	er := auth.authService.Register(&user)
	if er != nil {
		reponse.ReponseOutput(c, reponse.RegisterFail, "", nil)
		return
	}
	user.Password = ""
	reponse.ReponseOutput(c, reponse.RegisterSucess, "", user)
}

func (auth *AuthHandler) VerifyOTP(c *gin.Context) {

	var input request.OTPVerificationRequest

	if er := c.ShouldBind(&input); er != nil {
		reponse.ReponseOutput(c, reponse.OTPVerifyFail, "", nil)
		return
	}
	name, er := auth.authService.VerificationOTP(input.Email, input.OTP)
	if er != nil {
		reponse.ReponseOutput(c, reponse.OTPVerifyFail, "", nil)
		return
	}
	reponse.ReponseOutput(c, reponse.OTPVerifySucess, "", model.User{Email: input.Email, Name: name})
}

func (auth *AuthHandler) ResendOTP(c *gin.Context) {
	var input request.ResendOTPRequest
	if er := c.ShouldBind(&input); er != nil {
		reponse.ReponseOutput(c, reponse.ResendOTPFail, "", nil)
		return
	}

	_, er := auth.authService.ResendOTP(input.Email)
	if er != nil {
		reponse.ReponseOutput(c, reponse.ResendOTPFail, "", nil)
		return
	}
	reponse.ReponseOutput(c, reponse.ResendOTPSucess, "", model.User{Email: input.Email})
}
