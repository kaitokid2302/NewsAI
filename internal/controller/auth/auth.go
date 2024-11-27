package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/service"
)

type AuthController struct {
	emailService service.EmailService
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

type RegisterResponse struct {
	statusCode int
	data       database.User
	er         string
	message    string
}

func (auth *AuthController) Register(c *gin.Context) {
	var user database.User
	var response RegisterResponse
	if er := c.ShouldBind(&user); er != nil {
		response = RegisterResponse{
			statusCode: http.StatusBadRequest,
			er:         er.Error(),
			data:       database.User{},
			message:    "Invalid request",
		}
		c.JSON(response.statusCode, response)
		return
	}

	er := auth.emailService.SendEmail(user.Email)
	if er != nil {
		response = RegisterResponse{
			statusCode: http.StatusInternalServerError,
			er:         er.Error(),
			data:       database.User{},
			message:    "Internal server error",
		}
		c.JSON(response.statusCode, response)
		return
	}
	response = RegisterResponse{
		statusCode: http.StatusOK,
		message:    "User registered successfully. An OTP has been sent to your email.",
		data:       user,
		er:         "",
	}
	c.JSON(response.statusCode, response)
}
