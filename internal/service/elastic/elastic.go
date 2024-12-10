package elastic

import (
	"encoding/json"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/elastic"
)

type ElasticService interface {
	InsertToIndex(e *elastic.ElasticModel) error
	GetTextFromIndex(text string, from int, size int) ([]uint, error)
	AddSummaryToIndex(articleID uint, summary string) error
	FindDocument(articleID uint) (*elastic.ElasticModel, error)
	DeleteDocument(articleID uint) error
}

type elasticService struct {
	client *elasticsearch.Client
}

func NewElasticService(client *elasticsearch.Client) ElasticService {
	return &elasticService{client: client}
}

func (s *elasticService) InsertToIndex(e *elastic.ElasticModel) error {
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

func (s *elasticService) DeleteDocument(articleID uint) error {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"articleID": articleID,
			},
		},
	}
	queryByte, er := json.Marshal(query)
	if er != nil {
		return er
	}
	queryString := string(queryByte)
	_, err := s.client.DeleteByQuery(
		[]string{config.Global.IndexName},
		strings.NewReader(queryString),
		s.client.DeleteByQuery.WithRefresh(true),
	)
	if err != nil {
		return err
	}
	return nil
}

func (s *elasticService) FindDocument(articleID uint) (*elastic.ElasticModel, error) {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"articleID": articleID,
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
	if len(r["hits"].(map[string]interface{})["hits"].([]interface{})) == 0 {
		return nil, nil
	}
	hit := r["hits"].(map[string]interface{})["hits"].([]interface{})[0]
	source := hit.(map[string]interface{})["_source"].(map[string]interface{})
	var article *elastic.ElasticModel
	sourceByte, er := json.Marshal(source)
	if er != nil {
		return nil, er
	}
	er = json.Unmarshal(sourceByte, &article)
	if er != nil {
		return nil, er
	}
	return article, nil
}

func (s *elasticService) AddSummaryToIndex(articleID uint, summary string) error {
	// find the document
	// delete the document
	// insert the document with new summary
	article, er := s.FindDocument(articleID)
	if er != nil {
		return er
	}
	if article == nil {
		return nil
	}
	er = s.DeleteDocument(articleID)
	if er != nil {
		return er
	}
	article.Summary = summary
	er = s.InsertToIndex(article)
	if er != nil {
		return er
	}
	return nil
}

func (s *elasticService) GetTextFromIndex(text string, from int, size int) ([]uint, error) {
	// find on both *elasticsearch.ElasticModel.Text and *elasticsearch.ElasticModel.Summary
	query := map[string]interface{}{
		"from": from,
		"size": size,
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
