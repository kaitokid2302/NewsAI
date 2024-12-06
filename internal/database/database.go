package database

import (
	"fmt"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Global.Database.Host, config.Global.Database.User, config.Global.Database.Password, config.Global.Database.Database, config.Global.Database.Port)
	db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		panic(er)
	}
	err := db.AutoMigrate(&model.Topic{}, &model.User{}, &model.Article{})
	if err != nil {
		return nil
	}
	InitTopic(db)
	return db
}

func InitTopic(db *gorm.DB) {
	// model.topics
	er := db.Debug().Save(&model.Topics).Error
	if er != nil {
		panic(er)
	}
}
