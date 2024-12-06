package elasticsearch

import "github.com/elastic/go-elasticsearch/v8"

func InitElasticSearch() *elasticsearch.Client {
	client, er := elasticsearch.NewClient(elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	})
	if er != nil {
		panic(er)
	}
	return client
}
