package jwt

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/kaitokid2302/NewsAI/internal/config"
)

type JWTservice interface {
	CreateToken(email string) string
	VerifyToken(tokenIn string) (bool, string)
}

type JWTServiceImpl struct{}

func (x *JWTServiceImpl) CreateToken(email string) string {
	var key = config.Global.Key
	var t *jwt.Token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":   "mra2322001",
		"role":  "user",
		"exp":   time.Now().Add(time.Hour * 24 * 60).Unix(),
		"email": email,
	})
	s, _ := t.SignedString([]byte(key))
	return s
}

func (x *JWTServiceImpl) VerifyToken(tokenIn string) (bool, string) {
	key := config.Global.Key
	token, er := jwt.Parse(tokenIn, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("can not decode")
		}
		return []byte(key), nil
	})
	if er != nil {
		return false, ""
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return true, fmt.Sprintf("%v", claims["email"])
	} else {
		return false, ""
	}

}

func NewJWTService() JWTservice {
	return &JWTServiceImpl{}
}
