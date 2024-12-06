package repository

import (
	"github.com/kaitokid2302/NewsAI/internal/database"
	"gorm.io/gorm"
)

type UserRepo interface {
	SaveUserDB(user *database.User) error
	GetUserByEmail(email string) (*database.User, error)
	ExistUser(email string) bool
	Login(email, password string) (*database.User, error)
}

type UserRepoImpl struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &UserRepoImpl{db: db}
}

func (repo *UserRepoImpl) Login(email, password string) (*database.User, error) {
	var user database.User
	repo.db.Debug().Where("email = ? AND password = ?", email, password).First(&user)
	if user.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}
	user.Password = ""
	return &user, nil

}

func (repo *UserRepoImpl) SaveUserDB(user *database.User) error {
	db := repo.db
	return db.Debug().Save(user).Error
}

func (repo *UserRepoImpl) GetUserByEmail(email string) (*database.User, error) {
	db := repo.db
	var user database.User
	if err := db.Debug().Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepoImpl) ExistUser(email string) bool {
	db := repo.db
	count := db.Debug().Where("email = ?", email).Find(&database.User{}).RowsAffected
	return count > 0
}
