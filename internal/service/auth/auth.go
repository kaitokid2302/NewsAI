package auth

import (
	"errors"

	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository/user"
	"github.com/redis/go-redis/v9"
)

type AuthService interface {
	Register(user *database.User) error
	VerificationOTP(email string, code int) (string, error)
	SendEmail(to string) (int, error)
	SetOTPCode(email string, code int) error
	GetOTPCode(email string) (int, error)
	ResendOTP(email string) (int, error)
	Login(email, password string) (*database.User, error)
}

type AuthServiceImpl struct {
	userRepo    user.UserRepo
	redisClient *redis.Client
}

func NewAuthService(repo user.UserRepo, redisClient *redis.Client) AuthService {
	return &AuthServiceImpl{userRepo: repo, redisClient: redisClient}
}

func (s *AuthServiceImpl) Login(email, password string) (*database.User, error) {
	user, er := s.userRepo.Login(email, password)
	if er != nil {
		return nil, er
	}
	user.Password = ""
	return user, nil
}

func (s *AuthServiceImpl) Register(user *database.User) error {
	exist := s.userRepo.ExistUser(user.Email)
	if exist {
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

func (s *AuthServiceImpl) VerificationOTP(email string, code int) (string, error) {
	exist := s.userRepo.ExistUser(email)
	if exist {
		return "", errors.New("user already exist")
	}
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
