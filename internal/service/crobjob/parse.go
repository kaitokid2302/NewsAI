package crobjob

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
)

type Rss struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Image       Image  `xml:"image"`
	PubDate     string `xml:"pubDate"`
	Generator   string `xml:"generator"`
	Link        string `xml:"link"`
	Items       []Item `xml:"item"`
}

type Image struct {
	URL   string `xml:"url"`
	Title string `xml:"title"`
	Link  string `xml:"link"`
}

type Item struct {
	Title       string    `xml:"title"`
	Description string    `xml:"description"`
	PubDate     string    `xml:"pubDate"`
	Link        string    `xml:"link"`
	Guid        string    `xml:"guid"`
	Enclosure   Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	Type   string `xml:"type,attr"`
	Length string `xml:"length,attr"`
	URL    string `xml:"url,attr"`
}

func (r *CronJobArticleService) ArticleFromTopic(topic *database.Topic) ([]*database.Article, error) {
	link := topic.RssLink
	res, er := http.Get(link)
	if er != nil {
		return nil, er
	}
	defer res.Body.Close()
	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		return nil, er
	}
	var rss Rss
	er = xml.Unmarshal(data, &rss)
	if er != nil {
		return nil, er
	}
	var articles []*database.Article
	for _, item := range rss.Channel.Items {
		// description, after /br>
		description := item.Description
		// find index of </br>
		index := 0
		for i := len(description) - 5; i >= 0; i-- {
			if description[i:i+5] == "</br>" {
				index = i
				break
			}
		}
		// get description after </br>
		item.Description = description[index+5:]
		article := database.Article{
			Title:          item.Title,
			ImageEnclosure: item.Enclosure.URL,
			Description:    item.Description,
			PubDate:        item.PubDate,
			Link:           item.Link,
			TopicID:        topic.ID,
		}
		articles = append(articles, &article)
	}
	return articles, nil
}
