package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/handler/auth"
	"github.com/kaitokid2302/NewsAI/internal/service/jwt"
)

type Auth struct {
	jwtService jwt.JWTservice
}

func NewAuth(jwtService jwt.JWTservice) *Auth {
	return &Auth{
		jwtService: jwtService,
	}
}

func (a *Auth) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		var reponse auth.Response
		if token == "" {
			reponse = auth.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "token not correct",
			}

			c.JSON(http.StatusUnauthorized, reponse)
			c.Abort()
			return
		}

		ok, email := a.jwtService.VerifyToken(token)
		if !ok {
			reponse = auth.Response{
				StatusCode: http.StatusUnauthorized,
				Message:    "token not correct",
			}
			c.JSON(http.StatusUnauthorized, reponse)
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Next()
	}
}
