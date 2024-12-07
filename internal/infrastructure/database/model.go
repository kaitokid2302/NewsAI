package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `form:"name" binding:"required" json:"name,omitempty"`
	// email validate
	Email         string  `gorm:"unique" binding:"required,email" form:"email" json:"email,omitempty"`
	Password      string  `validate:"binding" form:"password" json:"password,omitempty"`
	Avatar        string  `form:"avatar" json:"avatar,omitempty"`
	TopicInterest []Topic `gorm:"many2many:user_topics" json:"topic_interest,omitempty"`
}

type Topic struct {
	gorm.Model
	Name    string `form:"name" binding:"required" json:"name,omitempty"`
	RssLink string `form:"rss_link" binding:"required" json:"rss_link,omitempty"`
}

// <item>
// <title>Giá won lao dốc khi Tổng thống Hàn Quốc ban bố thiết quân luật</title>
// <description>
// <![CDATA[ <a href="https://vnexpress.net/gia-won-lao-doc-khi-tong-thong-han-quoc-ban-bo-thiet-quan-luat-4823455.html"><img src="https://vcdn1-kinhdoanh.vnecdn.net/2024/12/03/krw-1733241822-1449-1733241835.jpg?w=1200&h=0&q=100&dpr=1&fit=crop&s=C3L9-MWeWL7USdxtsr4U5w"></a></br>Đồng won mất giá 2,5% so với đôla Mỹ, xuống thấp nhất kể từ năm 2016 sau sắc lệnh của Tổng thống Yoon Suk-yeol. ]]>
// </description>
// <pubDate>Tue, 03 Dec 2024 23:07:48 +0700</pubDate>
// <link>https://vnexpress.net/gia-won-lao-doc-khi-tong-thong-han-quoc-ban-bo-thiet-quan-luat-4823455.html</link>
// <guid>https://vnexpress.net/gia-won-lao-doc-khi-tong-thong-han-quoc-ban-bo-thiet-quan-luat-4823455.html</guid>
// <enclosure type="image/jpeg" length="1200" url="https://vcdn1-kinhdoanh.vnecdn.net/2024/12/03/krw-1733241822-1449-1733241835.jpg?w=1200&h=0&q=100&dpr=1&fit=crop&s=C3L9-MWeWL7USdxtsr4U5w"/>
// </item>
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
