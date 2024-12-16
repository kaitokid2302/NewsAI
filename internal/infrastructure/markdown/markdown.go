package markdown

import (
	"fmt"
	"net/http"

	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
)

type MarkdownInfrast interface {
	GetMarkDownFromLink(title, description, link string) (string, error)
}

func NewMarkdown() MarkdownInfrast {
	return &markdownInfrastImpl{}
}

type markdownInfrastImpl struct {
}

func (m *markdownInfrastImpl) GetMarkDownFromLink(title, description, link string) (string, error) {
	url := link
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	// Tạo converter với rules tùy chỉnh
	converter := md.NewConverter("", true, nil)

	// Thêm rule xử lý ảnh
	converter.AddRules(md.Rule{
		Filter: []string{"img"},
		Replacement: func(content string, selec *goquery.Selection, opt *md.Options) *string {
			// Lấy alt text
			alt := selec.AttrOr("alt", "")

			// Ưu tiên lấy src từ data-src
			src, exists := selec.Attr("data-src")
			if !exists || src == "" {
				src = selec.AttrOr("src", "")
			}

			// Bỏ qua ảnh placeholder hoặc ảnh base64
			if src == "" || src == "data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw==" {
				return nil
			}

			markdown := fmt.Sprintf("![%s](%s)", alt, src)
			return &markdown
		},
	})

	// Lấy nội dung từ article.fck_detail
	content, _ := doc.Find("article.fck_detail").Html()
	markdown, err := converter.ConvertString(content)
	if err != nil {
		return "", err
	}

	s := fmt.Sprintf("%s %s\n\n %s\n\n%s", "#", title, description, markdown)
	return s, nil
}
