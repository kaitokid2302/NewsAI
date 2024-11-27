package auth

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/service"
	"github.com/redis/go-redis/v9"
)

type AuthHandler struct {
	emailService service.EmailService
	redisClient  *redis.Client
}

func NewAuthHandler(emailService service.EmailService, redisClient *redis.Client) *AuthHandler {
	return &AuthHandler{
		emailService,
		redisClient,
	}
}

type RegisterResponse struct {
	statusCode int
	data       database.User
	er         string
	message    string
}

func (auth *AuthHandler) Register(c *gin.Context) {
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

	otpCode, er := auth.emailService.SendEmail(user.Email)
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
		message:    "User registered successfully. An OTP has been sent to your email and is valid for 5 minutes.",
		data:       user,
		er:         "",
	}

	auth.redisClient.SetEx(c, user.Email, otpCode, time.Minute*5)
	c.JSON(response.statusCode, response)
}
