package repository

import (
	"github.com/kaitokid2302/NewsAI/internal/database"
	"gorm.io/gorm"
)

type ArticleRepo interface {
	SaveArticle(article *database.Article) error
	ExistArticleByLink(link string) bool
}

type articleRepo struct {
	db *gorm.DB
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
