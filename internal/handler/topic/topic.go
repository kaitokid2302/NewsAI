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
	er := h.topicService.Subscribe(topicName)
	if er != nil {
		reponse.ReponseOutput(c, reponse.SubscribeTopicFail, "", nil)
	}
}
