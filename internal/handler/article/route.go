package article

import "github.com/gin-gonic/gin"

func (h *ArticleHandler) InitRoute(r *gin.RouterGroup) {
	r.GET("/article/:articleID", h.GetArticle)
	r.GET("/article", h.AllArticle)
	r.PUT("/article/viewed/:articleID", h.MarkViewed)
	r.PUT("/article/bookmark/:articleID", h.MarkBookMark)
	r.PUT("/article/hidden/:articleID", h.MarkHidden)
	r.DELETE("/article/viewed/:articleID", h.UnMarkViewed)
	r.DELETE("/article/bookmark/:articleID", h.UnMarkBookMark)
	r.DELETE("/article/hidden/:articleID", h.UnMarkHidden)

	// get text markdown from article

	r.GET("/article/text/:articleID", h.GetTextFromArticle)
}
