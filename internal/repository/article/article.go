package article

import (
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	SaveArticle(article *database.Article) error
	ExistArticleByLink(link string) bool
	GetArticle(articleID int) (*database.Article, error)
	ViewedArticle(userID int, offset int, limit int) (*[]database.Article, error)
	HiddenArticle(userID int, offset int, limit int) (*[]database.Article, error)
	BookMarkArticle(userID int, offset int, limit int) (*[]database.Article, error)
	ExistViewedArticle(userID int, articleID int) (bool, error)
	InsertToViewTable(userID int, articleID int) error
	ExistHiddenArticle(userID int, articleID int) (bool, error)
	ExistBookMarkArticle(userID int, articleID int) (bool, error)
	InsertToBookMarkTable(userID int, articleID int) error
	RemoveViewedArticle(userID int, articleID int) error
	RemoveHiddenArticle(userID int, articleID int) error
	RemoveBookMarkArticle(userID int, articleID int) error
	InsertToHiddenTable(userID int, articleID int) error
}

type articleRepo struct {
	db *gorm.DB
}

func (a *articleRepo) GetArticle(articleID int) (*database.Article, error) {
	var article database.Article
	er := a.db.First(&article, articleID).Error
	return &article, er
}

func NewArticleRepo(db *gorm.DB) ArticleRepo {
	return &articleRepo{db: db}
}

func (a *articleRepo) ExistArticleByLink(link string) bool {
	var article database.Article
	count := a.db.Where("link = ?", link).First(&article).RowsAffected
	return count > 0
}

func (a *articleRepo) SaveArticle(article *database.Article) error {
	if a.ExistArticleByLink(article.Link) {
		return nil
	}
	return a.db.Save(article).Error
}

func (a *articleRepo) ViewedArticle(userID int, offset int, limit int) (*[]database.Article, error) {
	var user database.User
	er := a.db.Debug().Preload("ViewedArticle").Offset(offset).Limit(limit).First(&user, userID).Error
	if er != nil {
		return nil, er
	}
	articles := user.ViewedArticle
	return &articles, nil
}

func (a *articleRepo) HiddenArticle(userID int, offset int, limit int) (*[]database.Article, error) {
	var user database.User
	er := a.db.Debug().Preload("HiddenArticle").Offset(offset).Limit(limit).First(&user, userID).Error
	if er != nil {
		return nil, er
	}
	articles := user.HiddenArticle
	return &articles, nil
}

func (a *articleRepo) BookMarkArticle(userID int, offset int, limit int) (*[]database.Article, error) {
	var user database.User
	er := a.db.Debug().Preload("BookMark").Offset(offset).Limit(limit).First(&user, userID).Error
	if er != nil {
		return nil, er
	}
	articles := user.BookMark
	return &articles, nil
}

func (a *articleRepo) ExistViewedArticle(userID int, articleID int) (bool, error) {
	// check table user_viewed_articles
	var count int64
	er := a.db.Debug().Table("user_viewed_articles").Where("user_id = ? AND article_id = ?", userID, articleID).Count(&count).Error
	return count > 0, er
}

func (a *articleRepo) InsertToViewTable(userID int, articleID int) error {
	// save to table user_viewed_articles
	er := a.db.Debug().Table("user_viewed_articles").Create(map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	}).Error
	return er
}

func (a *articleRepo) ExistHiddenArticle(userID int, articleID int) (bool, error) {
	// check table user_hidden_articles
	var count int64
	er := a.db.Debug().Table("user_hidden_articles").Where("user_id = ? AND article_id = ?", userID, articleID).Count(&count).Error
	return count > 0, er
}

func (a *articleRepo) ExistBookMarkArticle(userID int, articleID int) (bool, error) {
	// check table user_bookmarks
	var count int64
	er := a.db.Debug().Table("user_bookmarks").Where("user_id = ? AND article_id = ?", userID, articleID).Count(&count).Error
	return count > 0, er
}

func (a *articleRepo) InsertToBookMarkTable(userID int, articleID int) error {
	// save to table user_bookmarks
	er := a.db.Debug().Table("user_bookmarks").Create(map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	}).Error
	return er
}

func (a *articleRepo) RemoveViewedArticle(userID int, articleID int) error {
	var count int64
	er := a.db.Debug().Table("user_viewed_articles").Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&count).Error
	if er != nil {
		return er
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (a *articleRepo) RemoveHiddenArticle(userID int, articleID int) error {
	var count int64
	er := a.db.Debug().Table("user_hidden_articles").Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&count).Error
	if er != nil {
		return er
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (a *articleRepo) RemoveBookMarkArticle(userID int, articleID int) error {
	var count int64
	er := a.db.Debug().Table("user_bookmarks").Where("user_id = ? AND article_id = ?", userID, articleID).Delete(&count).Error
	if er != nil {
		return er
	}
	if count == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

func (a *articleRepo) InsertToHiddenTable(userID int, articleID int) error {
	er := a.db.Debug().Table("user_hidden_articles").Create(map[string]interface{}{
		"user_id":    userID,
		"article_id": articleID,
	}).Error
	return er
}
