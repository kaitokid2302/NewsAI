package database

import (
	"fmt"

	. "github.com/kaitokid2302/NewsAI/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", Global.Host, Global.User, Global.Password, Global.Database.Database, Global.Port)
	db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		panic(er)
	}
	db.AutoMigrate(&User{})
	return db
}
