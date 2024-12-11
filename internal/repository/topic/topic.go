package topic

import (
	"errors"

	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"gorm.io/gorm"
)

type TopicRepository interface {
	Subscribe(userID uint, topicID uint) error
	Unsubscribe(userID uint, topicID uint) error
	AllTopicOfUser(userID uint) (*[]database.Topic, error)
	FindTopicByName(topicName string) (*database.Topic, error)
}

type topicRepositoryImpl struct {
	db *gorm.DB
}

func NewTopicRepository(db *gorm.DB) TopicRepository {
	return &topicRepositoryImpl{
		db: db,
	}
}

func (r *topicRepositoryImpl) FindTopicByName(topicName string) (*database.Topic, error) {
	var topic database.Topic
	er := r.db.Debug().Where("name = ?", topicName).First(&topic).Error
	if er != nil {
		return nil, er
	}
	if topic.ID == 0 {
		return nil, errors.New("topic not exist")
	}
	return &topic, nil
}

func (r *topicRepositoryImpl) Subscribe(userID uint, topicID uint) error {
	// find userid, topicid in user_topics
	// if not exist, create new record

	var user database.User
	er := r.db.Debug().Preload("TopicInterest").First(&user, userID).Error
	if er != nil {
		return er
	}
	if user.ID == 0 {
		return errors.New("user not exist")
	}
	topicInterest := user.TopicInterest
	var found = false
	for i := 0; i < len(topicInterest); i++ {
		if topicInterest[i].ID == topicID {
			found = true
		}
	}

	if found {
		return errors.New("user already subscribed to this topic")
	}

	topic := database.Topic{}
	topic.ID = topicID
	user.TopicInterest = append(user.TopicInterest, topic)

	er = r.db.Debug().Save(&user).Error
	if er != nil {
		return er
	}
	return nil
}

func (r *topicRepositoryImpl) Unsubscribe(userID uint, topicID uint) error {

	var user database.User

	er := r.db.Debug().Preload("TopicInterest").First(&user, userID).Error
	if er != nil {
		return er
	}

	if user.ID == 0 {
		return errors.New("user not exist")
	}

	var topicToDelete database.Topic
	var found = false
	for i := 0; i < len(user.TopicInterest); i++ {
		if user.TopicInterest[i].ID == topicID {
			found = true
			topicToDelete.ID = user.TopicInterest[i].ID
			break
		}
	}
	if !found {
		return errors.New("user not subscribed to this topic")
	}
	er = r.db.Debug().Model(&user).Association("TopicInterest").Delete(&topicToDelete)
	return er
}

func (r *topicRepositoryImpl) AllTopicOfUser(userID uint) (*[]database.Topic, error) {
	var user database.User
	er := r.db.Debug().Preload("TopicInterest").First(&user, userID).Error
	if er != nil {
		return nil, er
	}
	if user.ID == 0 {
		return nil, errors.New("user not exist")
	}
	return &user.TopicInterest, nil
}
