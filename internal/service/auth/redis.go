package auth

import (
	"context"
	"encoding/json"
	"strconv"
	"time"

	"github.com/kaitokid2302/NewsAI/internal/database"
)

func (s *UserServiceImpl) SetOTPCode(email string, code int) error {
	er := s.redisClient.SetEx(context.Background(), email, code, time.Minute*5).Err()
	return er
}

func (s *UserServiceImpl) GetOTPCode(email string) (int, error) {
	code := s.redisClient.Get(context.Background(), email).Val()
	intCode, er := strconv.Atoi(code)
	return intCode, er
}

func (s *UserServiceImpl) SaveTempUser(user *database.User) error {
	x, er := json.Marshal(user)
	if er != nil {
		return er
	}
	er = s.redisClient.SetEx(context.Background(), user.Email+"temp", string(x), time.Hour).Err()
	return er
}

func (s *UserServiceImpl) GetTempUser(email string) (*database.User, error) {
	x := s.redisClient.Get(context.Background(), email+"temp").Val()
	var user database.User
	er := json.Unmarshal([]byte(x), &user)
	return &user, er
}
