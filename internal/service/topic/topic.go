package topic

import (
	"errors"

	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository/topic"
	"github.com/kaitokid2302/NewsAI/internal/service/user"
)

type TopicService interface {
	Subscribe(email string, topic string) error
	Unsubscribe(email string, topic string) error
	AllTopic(email string) (*[]database.Topic, error)
}

type topicServiceImpl struct {
	userService user.UserService
	topicRepo   topic.TopicRepository
}

func (s *topicServiceImpl) AllTopic(email string) (*[]database.Topic, error) {
	user, er := s.userService.GetUserInfo(email)
	if er != nil {
		return nil, er
	}
	topics, er := s.topicRepo.AllTopicOfUser(user.ID)
	if er != nil {
		return nil, er
	}
	return topics, nil
}

func NewTopicService(userService user.UserService, topicRepo topic.TopicRepository) TopicService {
	return &topicServiceImpl{
		userService: userService,
		topicRepo:   topicRepo,
	}
}

func (s *topicServiceImpl) Subscribe(email string, topic string) error {
	topicModel, er := s.topicRepo.FindTopicByName(topic)
	if er != nil {
		return er
	}
	if topicModel == nil {
		return errors.New("topic not exist")
	}
	userModel, er := s.userService.GetUserInfo(email)
	if er != nil {
		return er
	}
	er = s.topicRepo.Subscribe(userModel.ID, topicModel.ID)
	return er
}

func (s *topicServiceImpl) Unsubscribe(email string, topic string) error {
	topicModel, er := s.topicRepo.FindTopicByName(topic)
	if er != nil {
		return er
	}
	if topicModel == nil {
		return errors.New("topic not exist")
	}
	userModel, er := s.userService.GetUserInfo(email)
	if er != nil {
		return er
	}
	er = s.topicRepo.Unsubscribe(userModel.ID, topicModel.ID)
	return er
}
