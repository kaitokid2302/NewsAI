package article

import (
	"errors"

	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository/article"
	"github.com/kaitokid2302/NewsAI/internal/repository/topic"
	user2 "github.com/kaitokid2302/NewsAI/internal/repository/user"
	"github.com/kaitokid2302/NewsAI/internal/request"
	"gorm.io/gorm"
)

type ArticleService interface {
	GetArticle(articleID int) (*database.Article, error)
	AllArticle(email string, input *request.ArticleQueryRequest) (*[]database.Article, error)
	MarkViewed(email string, articleID int) error
	MarkBookMark(email string, id int) error
	MarkHidden(email string, id int) error
	UnMarkViewed(email string, id int) error
	UnMarkBookMark(email string, id int) error
	UnMarkHidden(email string, id int) error
}

type articleServiceImpl struct {
	articleRepo     article.ArticleRepo
	userRepository  user2.UserRepo
	topicRepository topic.TopicRepository
}

func NewArticleService(articleRepo article.ArticleRepo, userRepository user2.UserRepo, topicRepository topic.TopicRepository) ArticleService {
	return &articleServiceImpl{
		articleRepo:     articleRepo,
		userRepository:  userRepository,
		topicRepository: topicRepository,
	}
}

func (s *articleServiceImpl) GetArticle(articleID int) (*database.Article, error) {
	article, er := s.articleRepo.GetArticle(articleID)
	return article, er
}

func (s *articleServiceImpl) AllArticle(email string, input *request.ArticleQueryRequest) (*[]database.Article, error) {
	count := 0
	if input.Hidden {
		count++
	}
	if input.Viewed {
		count++
	}
	if input.BookMark {
		count++
	}
	if count > 1 {
		return nil, errors.New("only one of Hidden, Viewed, BookMark can be true")
	} else if count == 0 {
		return nil, errors.New("one of Hidden, Viewed, BookMark must be true")
	}
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return nil, er
	}
	user.Password = ""
	if input.Viewed {
		articles, er := s.articleRepo.ViewedArticle(int(user.ID), input.Offset, input.Limit)
		if er != nil {
			return nil, er
		}
		return articles, nil
	}
	if input.Hidden {
		article, er := s.articleRepo.HiddenArticle(int(user.ID), input.Offset, input.Limit)
		if er != nil {
			return nil, er
		}
		return article, nil
	}
	if input.BookMark {
		article, er := s.articleRepo.BookMarkArticle(int(user.ID), input.Offset, input.Limit)
		if er != nil {
			return nil, er
		}
		return article, nil
	}
	return nil, nil
}

func (s *articleServiceImpl) MarkViewed(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_viewed_articles contain user.ID and article.ID
	exist, er := s.articleRepo.ExistViewedArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist {
		return errors.New("article already viewed")
	}
	// check hidden table
	hidden, er := s.articleRepo.ExistHiddenArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if hidden {
		return errors.New("can not mark viewed because article is hidden")
	}
	// save to table user_viewed_articles
	return s.articleRepo.InsertToViewTable(int(user.ID), articleID)
}

func (s *articleServiceImpl) MarkBookMark(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_bookmarks contain user.ID and article.ID
	exist, er := s.articleRepo.ExistBookMarkArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist {
		return errors.New("article already bookmarked")
	}
	// hidden can not book mark, check hidden table
	hidden, er := s.articleRepo.ExistHiddenArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if hidden {
		return errors.New("can not mark bookmark because article is hidden")
	}
	// save to table user_bookmarks
	return s.articleRepo.InsertToBookMarkTable(int(user.ID), articleID)
}

func (s *articleServiceImpl) MarkHidden(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_hidden_articles contain user.ID and article.ID
	exist, er := s.articleRepo.ExistHiddenArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist {
		return errors.New("article already hidden")
	}
	// remove from user_viewed_articles
	// remove from user_bookmarks
	// save to table user_hidden_articles

	er = s.articleRepo.RemoveViewedArticle(int(user.ID), articleID)
	if errors.Is(er, gorm.ErrRecordNotFound) == false {
		return er
	}
	er = s.articleRepo.RemoveBookMarkArticle(int(user.ID), articleID)
	if errors.Is(er, gorm.ErrRecordNotFound) == false {
		return er
	}
	return s.articleRepo.InsertToHiddenTable(int(user.ID), articleID)
}

func (s *articleServiceImpl) UnMarkViewed(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_viewed_articles contain user.ID and article.ID
	exist, er := s.articleRepo.ExistViewedArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist == false {
		return errors.New("article not exist in viewed list")
	}
	// remove from user_viewed_articles
	return s.articleRepo.RemoveViewedArticle(int(user.ID), articleID)
}

func (s *articleServiceImpl) UnMarkBookMark(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_bookmarks contain user.ID and article.ID
	exist, er := s.articleRepo.ExistBookMarkArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist == false {
		return errors.New("article not exist in bookmark list")
	}
	// remove from user_bookmarks
	return s.articleRepo.RemoveBookMarkArticle(int(user.ID), articleID)
}

func (s *articleServiceImpl) UnMarkHidden(email string, articleID int) error {
	// put
	user, er := s.userRepository.GetUserByEmail(email)
	if er != nil {
		return er
	}
	article, er := s.articleRepo.GetArticle(articleID)
	if er != nil {
		return er
	}
	// check if this topicID of article == topicID of user
	topicInterest, er := s.topicRepository.AllTopicOfUser(user.ID)
	if er != nil {
		return er
	}
	topicAritcle := article.TopicID
	found := false
	for i := 0; i < len(*topicInterest); i++ {
		if (*topicInterest)[i].ID == topicAritcle {
			found = true
			break
		}
	}
	if found == false {
		return errors.New("user not subscribed to this topic")
	}
	// check table user_hidden_articles contain user.ID and article.ID
	exist, er := s.articleRepo.ExistHiddenArticle(int(user.ID), articleID)
	if er != nil {
		return er
	}
	if exist == false {
		return errors.New("article not exist in hidden list")
	}
	// remove from user_hidden_articles
	return s.articleRepo.RemoveHiddenArticle(int(user.ID), articleID)
}

// unViewed
// unBookMark
// unHidden
