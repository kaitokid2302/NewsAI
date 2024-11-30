package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
	userservice "github.com/kaitokid2302/NewsAI/internal/service/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
)

type AuthHandler struct {
	userService userservice.UserService
	jwtService  jwt.JWTservice
}

func NewAuthHandler(userService userservice.UserService, jwtService jwt.JWTservice) *AuthHandler {
	return &AuthHandler{
		userService: userService,
		jwtService:  jwtService,
	}
}

// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept multipart/form-data
// @Accept json
// @Produce json
// @Param   name     formData  string  true  "Username"
// @Param   email    formData  string  true  "Email"
// @Param   password formData  string  true  "Password"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} RegisterResponse
// @Failure 500 {object} RegisterResponse
// @Router /auth/register [post]
func (auth *AuthHandler) Register(c *gin.Context) {
	var user database.User
	var response Response
	if er := c.ShouldBind(&user); er != nil {
		response = Response{
			StatusCode: http.StatusBadRequest,
			Er:         er.Error(),
			Message:    "Invalid request",
		}
		c.JSON(response.StatusCode, response)
		return
	}
	er := auth.userService.Register(&user)
	if er != nil {
		response = Response{
			StatusCode: http.StatusInternalServerError,
			Er:         er.Error(),
			Message:    "Can not register user",
		}
		c.JSON(response.StatusCode, response)
		return
	}
	user.Password = ""
	response = Response{
		StatusCode: http.StatusOK,
		Data:       user,
		Message:    "OTP has been sent to your email. The code is only valid for 5 minutes.",
	}
	c.JSON(response.StatusCode, response)
}

// @Summary OTP authentication
// @Description OTP authentication
// @Tags auth
// @Accept json
// @Accept multipart/form-data
// @Produce json
// @Param request body OTPVerificationRequest true "OTP Verification Request"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} RegisterResponse
// @Failure 500 {object} RegisterResponse
// @Router /auth/verify [post]
func (auth *AuthHandler) VerifyOTP(c *gin.Context) {
	// email
	// otp

	var input OTPVerificationRequest

	var output Response

	if er := c.ShouldBind(&input); er != nil {
		output = Response{
			StatusCode: http.StatusBadRequest,
			Er:         er.Error(),
			Message:    "Invalid request",
		}
		c.JSON(output.StatusCode, output)
		return
	}
	name, er := auth.userService.VerificationOTP(input.Email, input.OTP)
	if er != nil {
		output = Response{
			StatusCode: http.StatusInternalServerError,
			Er:         er.Error(),
			Message:    "Can not verify OTP",
		}
		c.JSON(output.StatusCode, output)
		return
	}
	user := database.User{
		Email: input.Email,
		Name:  name,
	}
	output = Response{
		StatusCode: http.StatusOK,
		Message:    "User has been registered successfully",
		Data:       user,
	}
	c.JSON(output.StatusCode, output)
}

// @Summary Resend OTP
// @Description Resend OTP
// @Tags auth
// @Accept json
// @Accept multipart/form-data
// @Produce json
// @Param email formData string true "Email"
// @Success 200 {object} RegisterResponse
// @Failure 400 {object} RegisterResponse
// @Failure 500 {object} RegisterResponse
// @Router /auth/verify/resend [post]
func (auth *AuthHandler) ResendOTP(c *gin.Context) {
	var email struct {
		Email string `json:"email" binding:"required,email" form:"email"`
	}
	var output Response
	if er := c.ShouldBind(&email); er != nil {
		output = Response{
			StatusCode: http.StatusBadRequest,
			Er:         er.Error(),
			Message:    "Invalid request",
		}
		c.JSON(output.StatusCode, output)
		return
	}

	_, er := auth.userService.ResendOTP(email.Email)
	if er != nil {
		output = Response{
			StatusCode: http.StatusInternalServerError,
			Er:         er.Error(),
			Message:    "Can not resend OTP",
		}
		c.JSON(output.StatusCode, output)
		return
	}
	user := database.User{
		Email: email.Email,
	}
	output = Response{
		StatusCode: http.StatusOK,
		Message:    "OTP has been sent to your email. The code is only valid for 5 minutes.",
		Data:       user,
	}
	c.JSON(output.StatusCode, output)
}
