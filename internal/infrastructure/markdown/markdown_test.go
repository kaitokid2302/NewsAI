package markdown

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkDown(t *testing.T) {
	title := "Sự phối hợp giữa ông Biden và ông Trump phía sau lệnh ngừng bắn Israel - Hezbollah"
	description := "Khi chính quyền ông Biden thúc đẩy thỏa thuận ngừng bắn Israel - Hezbollah vào cuối nhiệm kỳ, họ có một đối tác bất ngờ: Tổng thống đắc cử Donald Trump."

	markdown := NewMarkdown()
	link := "https://vnexpress.net/su-phoi-hop-giua-ong-biden-va-ong-trump-phia-sau-lenh-ngung-ban-israel-hezbollah-4821718.html"

	ans, er := markdown.GetMarkDownFromLink(title, description, link)
	assert.Nil(t, er)
	print(ans)
}
