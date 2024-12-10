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
	data1 := elastic.ElasticModel{
		Text:      "trời ơi là trời",
		Summary:   "tôi phải làm sao đây",
		ArticleID: 100,
	}
	data2 := elastic.ElasticModel{
		Text:      "trời má là đất",
		Summary:   "em gái yêu",
		ArticleID: 101,
	}
	data3 := elastic.ElasticModel{
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

	dataList, er := elasticService.GetTextFromIndex("nhất ", 0, 3)
	assert.Nil(t, er)
	assert.Equal(t, 1, len(dataList))
	assert.Equal(t, uint(102), dataList[0])

	er = elasticService.DeleteDocument(102)
	assert.Nil(t, er)
	er = elasticService.DeleteDocument(101)
	assert.Nil(t, er)
	er = elasticService.DeleteDocument(100)
	assert.Nil(t, er)

	dataList, er = elasticService.GetTextFromIndex("nhất ", 0, 3)
	assert.Nil(t, er)
	assert.Equal(t, 0, len(dataList))
	// "trời"
	dataList, er = elasticService.GetTextFromIndex("trời", 0, 3)
	assert.Nil(t, er)
	assert.Equal(t, 0, len(dataList))
}
