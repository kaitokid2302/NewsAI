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
	StatusCode int                  `json:"statusCode"`
	Data       RegisterResponseData `json:"data"`
	Er         string               `json:"er"`
	Message    string               `json:"message"`
}

type RegisterResponseData struct {
	Email string `json:"email"`
	Name  string `json:"name"`
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
// @Router /register [post]
func (auth *AuthHandler) Register(c *gin.Context) {
	var user database.User
	var response RegisterResponse
	if er := c.ShouldBind(&user); er != nil {
		response = RegisterResponse{
			StatusCode: http.StatusBadRequest,
			Er:         er.Error(),
			Message:    "Invalid request",
		}
		c.JSON(response.StatusCode, response)
		return
	}

	otpCode, er := auth.emailService.SendEmail(user.Email)
	if er != nil {
		response = RegisterResponse{
			StatusCode: http.StatusInternalServerError,
			Er:         er.Error(),
			Message:    "Internal server error",
		}
		c.JSON(response.StatusCode, response)
		return
	}
	response = RegisterResponse{
		Data: RegisterResponseData{
			Email: user.Email,
			Name:  user.Name,
		},
		StatusCode: http.StatusOK,
		Message:    "User registered successfully. An OTP has been sent to your email and is valid for 5 minutes.",
		Er:         "",
	}

	auth.redisClient.SetEx(c, user.Email, otpCode, time.Minute*5)
	c.JSON(response.StatusCode, response)
}
