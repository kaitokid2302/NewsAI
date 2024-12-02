package user

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/database"
)

type UserReponse struct {
	StatusCode int           `json:"statusCode,omitempty"`
	Data       database.User `json:"data,omitempty"`
	Message    string        `json:"message,omitempty"`
}

func ErrorReponse(c *gin.Context, statusCode string, message string) {

}
