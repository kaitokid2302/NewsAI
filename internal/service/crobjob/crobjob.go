package crobjob

import (
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	elastic2 "github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/markdown"
	"github.com/kaitokid2302/NewsAI/internal/repository/article"
	"github.com/kaitokid2302/NewsAI/internal/service/elastic"
)

type CronJobArticleService struct {
	articleRepo    article.ArticleRepo
	markdown       markdown.Markdown
	elasticService elastic.ElasticService
}

func NewCronJobArticleService(articleRepo article.ArticleRepo, markdown markdown.Markdown, elasticService elastic.ElasticService) *CronJobArticleService {
	return &CronJobArticleService{
		articleRepo:    articleRepo,
		markdown:       markdown,
		elasticService: elasticService,
	}
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
		// save to elastic
		markdownText, er := c.markdown.GetMarkDownFromLink(article.Title, article.Description, article.Link)
		if er != nil {
			return er
		}
		err := c.elasticService.InsertToIndex(&elastic2.ElasticModel{
			Text:      markdownText,
			Summary:   "",
			ArticleID: article.ID,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
