package user

import (
	"mime/multipart"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	"github.com/kaitokid2302/NewsAI/internal/service/s3"
)

type UserService interface {
	UpdateUser(c *gin.Context, name string, fileName string, file *multipart.File) (*database.User, error)
	GetUserInfo(email string) (*database.User, error)
}

type UserServiceImpl struct {
	s3             s3.UploadFileS3Service
	userRepository repository.UserRepo
}

func NewUserService(s3 s3.UploadFileS3Service, userRepository repository.UserRepo) UserService {
	return &UserServiceImpl{
		s3,
		userRepository,
	}
}

func (u *UserServiceImpl) GetUserInfo(email string) (*database.User, error) {
	user, er := u.userRepository.GetUserByEmail(email)
	if er != nil {
		return nil, er
	}
	user.Password = ""
	return user, er
}

func (u *UserServiceImpl) UpdateUser(c *gin.Context, name string, fileName string, file *multipart.File) (*database.User, error) {
	user, er := u.userRepository.GetUserByEmail(c.GetString("email"))
	if er != nil {
		return nil, er
	}
	if file != nil {
		link, er := u.s3.UploadFile(fileName, *file)
		if er != nil {
			return nil, er
		}
		user.Avatar = link
	}
	user.Name = name
	er = u.userRepository.SaveUserDB(user)
	user.Password = ""
	return user, er
}
