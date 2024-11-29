package userservice

import (
	"errors"

	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	"github.com/redis/go-redis/v9"
)

type UserService interface {
	Register(user *database.User) error
	VerificationOTP(email string, code int) (string, error)
	SendEmail(to string) (int, error)
	SetOTPCode(email string, code int) error
	GetOTPCode(email string) (int, error)
}

type UserServiceImpl struct {
	userRepo    repository.UserRepo
	redisClient *redis.Client
}

func NewUserService(repo repository.UserRepo, redisClient *redis.Client) UserService {
	return &UserServiceImpl{userRepo: repo, redisClient: redisClient}
}

func (s *UserServiceImpl) Register(user *database.User) error {
	find, _ := s.userRepo.GetUserByEmail(user.Email)
	if find != nil {
		return errors.New("user already exists")
	}
	code, er := s.SendEmail(user.Email)
	if er != nil {
		return er
	}
	er = s.SetOTPCode(user.Email, code)
	if er != nil {
		return er
	}
	er = s.SaveTempUser(user)
	if er != nil {
		return er
	}
	return nil

}

func (s *UserServiceImpl) VerificationOTP(email string, code int) (string, error) {
	user, er := s.GetTempUser(email)
	if er != nil || user.Email == "" {
		return "", errors.New("user not found, register again")
	}
	otp, er := s.GetOTPCode(email)
	if er != nil {
		return "", er
	}
	if otp != code {
		return "", errors.New("invalid code")
	}
	er = s.userRepo.SaveUserDB(user)
	if er != nil {
		return "", er
	}
	return user.Name, nil
}
