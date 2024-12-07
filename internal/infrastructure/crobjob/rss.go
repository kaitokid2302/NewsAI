package crobjob

import (
	"time"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/service/crobjob"
)

type Crobjob struct {
	crobjobService crobjob.CronJobArticleService
}

func NewCrobjob(crobjobService crobjob.CronJobArticleService) *Crobjob {
	return &Crobjob{
		crobjobService: crobjobService,
	}
}

func (c *Crobjob) Run() {
	for {
		for _, topic := range database.Topics {
			c.crobjobService.GetArticle(&topic)
		}
		time.Sleep(time.Second * time.Duration(config.Global.Time))
	}
}
