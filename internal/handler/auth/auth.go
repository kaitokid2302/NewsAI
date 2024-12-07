package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	reponse2 "github.com/kaitokid2302/NewsAI/internal/reponse"
	"github.com/kaitokid2302/NewsAI/internal/request"
	authService "github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
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
		reponse2.ReponseOutput(c, reponse2.RegisterFail, "", nil)
		return
	}
	var user database.User
	automapper.MapLoose(registerRequest, &user)
	er := auth.authService.Register(&user)
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.RegisterFail, "", nil)
		return
	}
	user.Password = ""
	reponse2.ReponseOutput(c, reponse2.RegisterSucess, "", user)
}

func (auth *AuthHandler) VerifyOTP(c *gin.Context) {

	var input request.OTPVerificationRequest

	if er := c.ShouldBind(&input); er != nil {
		reponse2.ReponseOutput(c, reponse2.OTPVerifyFail, "", nil)
		return
	}
	name, er := auth.authService.VerificationOTP(input.Email, input.OTP)
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.OTPVerifyFail, "", nil)
		return
	}
	reponse2.ReponseOutput(c, reponse2.OTPVerifySucess, "", database.User{Email: input.Email, Name: name})
}

func (auth *AuthHandler) ResendOTP(c *gin.Context) {
	var input request.ResendOTPRequest
	if er := c.ShouldBind(&input); er != nil {
		reponse2.ReponseOutput(c, reponse2.ResendOTPFail, "", nil)
		return
	}

	_, er := auth.authService.ResendOTP(input.Email)
	if er != nil {
		reponse2.ReponseOutput(c, reponse2.ResendOTPFail, "", nil)
		return
	}
	reponse2.ReponseOutput(c, reponse2.ResendOTPSucess, "", database.User{Email: input.Email})
}
