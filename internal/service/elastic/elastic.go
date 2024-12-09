package elastic

import (
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
)

type ElasticService interface {
	InsertToIndex(e *elastic.Elastic) error
	SearchDocument(text string) ([]uint, error) // articleID
}

type elasticService struct {
	client *elasticsearch.Client
}

func NewElasticService(client *elasticsearch.Client) ElasticService {
	return &elasticService{client: client}
}

func (s *elasticService) InsertToIndex(e *elastic.Elastic) error {
	dataByte, er := json.Marshal(e)
	if er != nil {
		return er
	}
	dataString := string(dataByte)
	_, err := s.client.Index(
		config.Global.IndexName,
		strings.NewReader(dataString),
		s.client.Index.WithRefresh("true"),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *elasticService) SearchDocument(text string) ([]uint, error) {
	// find on both *elasticsearch.Elastic.Text and *elasticsearch.Elastic.Summary
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"should": []map[string]interface{}{
					{
						"match_phrase": map[string]interface{}{
							"text": map[string]interface{}{
								"query": text,
								"slop":  99999,
							},
						},
					},
					{
						"match_phrase": map[string]interface{}{
							"summary": map[string]interface{}{
								"query": text,
								"slop":  99999,
							},
						},
					},
				},
			},
		},
	}
	queryByte, er := json.Marshal(query)
	if er != nil {
		return nil, er
	}
	queryString := string(queryByte)
	res, err := s.client.Search(
		s.client.Search.WithIndex(config.Global.IndexName),
		s.client.Search.WithBody(strings.NewReader(queryString)),
	)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var r map[string]interface{}
	er = json.NewDecoder(res.Body).Decode(&r)
	if er != nil {
		return nil, er
	}

	var articleIDs []uint
	for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
		articleIDs = append(articleIDs, uint(hit.(map[string]interface{})["_source"].(map[string]interface{})["articleID"].(float64)))
	}
	return articleIDs, nil
}
