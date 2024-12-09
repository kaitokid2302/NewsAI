package elastic

import (
	"fmt"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/kaitokid2302/NewsAI/internal/config"
)

type Elastic struct {
	Text      string `json:"text"`
	Summary   string `json:"summary"`
	ArticleID uint   `json:"articleID"`
}

func InitElasticSearch() *elasticsearch.Client {
	client, er := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{fmt.Sprintf("http://%s:%d", config.Global.Elastic.Host, config.Global.Elastic.Port)},
	})
	if er != nil {
		panic(er)
	}
	exist, er := client.Indices.Exists([]string{config.Global.IndexName})
	if er != nil {
		panic(er)
	}
	if exist.StatusCode == 200 {
		return client
	}
	_, err := client.Indices.Create(config.Global.IndexName)
	if err != nil {
		panic(err)
	}
	return client
}
