package article

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kaitokid2302/NewsAI/internal/reponse"
	"github.com/kaitokid2302/NewsAI/internal/request"
	"github.com/kaitokid2302/NewsAI/internal/service/article"
)

type ArticleHandler struct {
	articleService article.ArticleService
}

func NewArticleHandler(articleService article.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: articleService}
}

func (h *ArticleHandler) GetArticle(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	article, er := h.articleService.GetArticle(articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", article)
}

func (h *ArticleHandler) AllArticle(c *gin.Context) {
	email := c.GetString("email")
	var input request.ArticleQueryRequest
	er := c.ShouldBind(&input)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	articles, er := h.articleService.AllArticle(email, &input)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", articles)
}

// markViewed
func (h *ArticleHandler) MarkViewed(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.MarkViewed(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}

// markBookMark

func (h *ArticleHandler) MarkBookMark(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.MarkBookMark(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}

// markHidden

func (h *ArticleHandler) MarkHidden(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.MarkHidden(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}

// unMarkViewed
func (h *ArticleHandler) UnMarkViewed(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.UnMarkViewed(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}

// unMarkBookMark
func (h *ArticleHandler) UnMarkBookMark(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.UnMarkBookMark(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}

// unMarkHidden
func (h *ArticleHandler) UnMarkHidden(c *gin.Context) {
	articleIDString := c.Param("articleID")
	articleID, er := strconv.Atoi(articleIDString)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	email := c.GetString("email")
	er = h.articleService.UnMarkHidden(email, articleID)
	if er != nil {
		reponse.ReponseOutput(c, reponse.GetArticleFail, er.Error(), nil)
		return
	}
	reponse.ReponseOutput(c, reponse.GetArticleSuccess, "", nil)
}
