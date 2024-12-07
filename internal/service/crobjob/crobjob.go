package crobjob

import (
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository"
)

type CronJobArticleService struct {
	articleRepo repository.ArticleRepo
}

func NewCronJobArticleService(articleRepo repository.ArticleRepo) *CronJobArticleService {
	return &CronJobArticleService{articleRepo: articleRepo}
}

func (c *CronJobArticleService) GetArticle(topic *database.Topic) error {
	articles, er := c.ArticleFromTopic(topic)
	if er != nil {
		return er
	}
	for _, article := range articles {
		er := c.articleRepo.SaveArticle(article)
		if er != nil {
			return er
		}
	}
	return nil
}
