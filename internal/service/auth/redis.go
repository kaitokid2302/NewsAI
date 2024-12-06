package auth

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/kaitokid2302/NewsAI/internal/database/model"
)

func (s *AuthServiceImpl) SetOTPCode(email string, code int) error {
	er := s.redisClient.SetEx(context.Background(), email, code, time.Minute*5).Err()
	return er
}

func (s *AuthServiceImpl) GetOTPCode(email string) (int, error) {
	code := s.redisClient.Get(context.Background(), email).Val()
	intCode, er := strconv.Atoi(code)
	return intCode, er
}

func (s *AuthServiceImpl) SaveTempUser(user *model.User) error {
	x, er := json.Marshal(user)
	if er != nil {
		return er
	}
	er = s.redisClient.SetEx(context.Background(), user.Email+"temp", string(x), time.Hour).Err()
	return er
}

func (s *AuthServiceImpl) GetTempUser(email string) (*model.User, error) {
	x := s.redisClient.Get(context.Background(), email+"temp").Val()
	var user model.User
	er := json.Unmarshal([]byte(x), &user)
	return &user, er
}
