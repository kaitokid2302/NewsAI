package topic

import "github.com/gin-gonic/gin"

func (h *TopicHandler) InitRoute(r *gin.RouterGroup) {
	r.PUT("/subscribe", h.Subscribe)
	r.PUT("/unsubscribe", h.Unsubscribe)
	r.GET("/all", h.AllTopic)
}
