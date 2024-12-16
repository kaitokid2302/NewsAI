package reponse

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ReponseStruct struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func ReponseOutput(c *gin.Context, code int, message string, data interface{}) {
	if message == "" {
		message = msg[code]
	}
	c.JSON(http.StatusOK, ReponseStruct{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
