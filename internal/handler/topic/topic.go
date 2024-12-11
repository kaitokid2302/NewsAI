package topic

import (
	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/reponse"
	"github.com/kaitokid2302/NewsAI/internal/service/topic"
)

type TopicHandler struct {
	topicService topic.TopicService
}

func (h *TopicHandler) Subscribe(c *gin.Context) {
	topicName := c.Query("topic_name")
	er := h.topicService.Subscribe(c.GetString("email"), topicName)
	if er != nil {
		reponse.ReponseOutput(c, reponse.SubscribeTopicFail, er.Error(), nil)
	}
	reponse.ReponseOutput(c, reponse.SubscribeTopicSuccess, "", nil)
}

func (h *TopicHandler) Unsubscribe(c *gin.Context) {
	topicName := c.Query("topic_name")
	er := h.topicService.Unsubscribe(c.GetString("email"), topicName)
	if er != nil {
		reponse.ReponseOutput(c, reponse.UnsubscribeTopicFail, er.Error(), nil)
	}
	reponse.ReponseOutput(c, reponse.UnsubscribeTopicSuccess, "", nil)
}
