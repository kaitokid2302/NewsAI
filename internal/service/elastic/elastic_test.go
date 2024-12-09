package elastic

import (
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
	"github.com/stretchr/testify/assert"
)

func TestNewElasticService(t *testing.T) {
	config.InitAll()
	client := elastic.InitElasticSearch()
	elasticService := NewElasticService(client)
	data1 := elastic.Elastic{
		Text:      "trời ơi là trời",
		Summary:   "tôi phải làm sao đây",
		ArticleID: 100,
	}
	data2 := elastic.Elastic{
		Text:      "trời má là đất",
		Summary:   "em gái yêu",
		ArticleID: 101,
	}
	data3 := elastic.Elastic{
		Text:      "con yêu mẹ",
		Summary:   "nhất nhà",
		ArticleID: 102,
	}
	er := elasticService.InsertToIndex(&data1)
	assert.Nil(t, er)
	er = elasticService.InsertToIndex(&data2)
	assert.Nil(t, er)
	er = elasticService.InsertToIndex(&data3)
	assert.Nil(t, er)

	dataList, er := elasticService.SearchDocument("trời là")
	assert.Nil(t, er)
	assert.Equal(t, 10, len(dataList))
}
