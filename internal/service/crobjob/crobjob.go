package crobjob

import (
	"github.com/kaitokid2302/NewsAI/internal/database"
	"github.com/kaitokid2302/NewsAI/internal/repository"
	"github.com/kaitokid2302/NewsAI/internal/rss"
)

type CronJobArticleService struct {
	articleRepo repository.ArticleRepo
	rssParser   rss.RssParser
}

func NewCronJobArticleService(articleRepo repository.ArticleRepo) *CronJobArticleService {
	return &CronJobArticleService{articleRepo: articleRepo}
}

func (c *CronJobArticleService) GetArticle(topic *database.Topic) error {
	articles, er := c.rssParser.GetArticle(topic)
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
