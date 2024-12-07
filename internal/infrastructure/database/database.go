package database

import (
	"fmt"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Global.Database.Host, config.Global.Database.User, config.Global.Database.Password, config.Global.Database.Database, config.Global.Database.Port)
	db, er := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if er != nil {
		panic(er)
	}
	err := db.AutoMigrate(&Topic{}, &User{}, &Article{})
	if err != nil {
		return nil
	}
	InitTopic(db)
	return db
}

func InitTopic(db *gorm.DB) {
	// model.topics
	for i := 0; i < len(Topics); i++ {
		t := &Topics[i]
		count := db.Where("name = ?", t.Name).Find(&Topic{}).RowsAffected
		if count == 0 {
			er := db.Save(t)
			if er.Error != nil {
				panic(er)
			}
		} else {
			er := db.Where("name = ?", t.Name).First(&Topics[i])
			if er.Error != nil {
				panic(er)
			}
		}
	}
}
