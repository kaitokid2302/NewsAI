package database

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string `form:"name" binding:"required" json:"name,omitempty"`
	// email validate
	Email         string    `gorm:"unique" binding:"required,email" form:"email" json:"email,omitempty"`
	Password      string    `validate:"binding" form:"password" json:"password,omitempty"`
	Avatar        string    `form:"avatar" json:"avatar,omitempty"`
	TopicInterest []Topic   `gorm:"many2many:user_topics" json:"topic_interest,omitempty"`
	BookMark      []Article `gorm:"many2many:user_bookmarks" json:"book_mark,omitempty"`
	ViewedArticle []Article `gorm:"many2many:user_viewed_articles" json:"viewed_article,omitempty"`
	HiddenArticle []Article `gorm:"many2many:user_hidden_articles" json:"hidden_article,omitempty"`
}

type Topic struct {
	gorm.Model
	Name    string `form:"name" binding:"required" json:"name,omitempty"`
	RssLink string `form:"rss_link" binding:"required" json:"rss_link,omitempty"`
}

type Article struct {
	gorm.Model
	Title          string `json:"title,omitempty"`
	ImageEnclosure string `json:"image,omitempty"`
	Description    string `json:"description,omitempty"`
	PubDate        string `json:"pubDate,omitempty"`
	Link           string `json:"link,omitempty"`
	TopicID        uint   `json:"topic_id,omitempty"`
	Topic          *Topic `json:"topic,omitempty"`
}
