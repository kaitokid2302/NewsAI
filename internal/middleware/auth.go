package middleware

import (
	"github.com/gin-gonic/gin"
	reponse2 "github.com/kaitokid2302/NewsAI/internal/reponse"
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

func (a *Auth) JWTverify() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			reponse2.ReponseOutput(c, reponse2.JWTVerifyFail, "", nil)
			c.Abort()
			return
		}

		ok, email := a.jwtService.VerifyToken(token)
		if !ok {
			reponse2.ReponseOutput(c, reponse2.JWTVerifyFail, "", nil)
			c.Abort()
			return
		}
		c.Set("email", email)
		c.Next()
	}
}
