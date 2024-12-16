package ai

import (
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAIInfrast(t *testing.T) {
	config.InitAll()
	provider := config.Global.AI.Provider[0]

	aiService := NewAIService(provider)

	ans, err := aiService.Summarize("-   [Thể thao](https://vnexpress.net/the-thao \"Thể thao\")\n-   [Bóng đá](https://vnexpress.net/bong-da \"Bóng đá\")\n-   [La Liga](https://vnexpress.net/bong-da/la-liga \"La Liga\")\n\nThứ hai, 16/12/2024, 06:43 (GMT+7)\n\nTây Ban NhaĐội dẫn đầu La Liga thua đối thủ xếp thứ 17, Leganes với tỷ số 0-1 ngay trên sân nhà, tối 15/12.\n\nBarca nhập cuộc thiếu tập trung, dẫn đến hậu quả là bàn thua ngay phút thứ 4. Sergio Gonzalez, trung vệ chơi 13 năm tại các hạng dưới trước khi cùng Leganes lên La Liga mùa này, thoải mái đánh đầu sau quả phạt góc đem lại ba điểm cho đội khách.\n\nBàn thua sớm khiến đội dẫn đầu La Liga phải đẩy mạnh tấn công. Những phút còn lại, các cầu thủ của HLV Hansi Flick kiểm soát bóng 81%, dứt điểm 20 lần và trúng đích 4 lần. Nhưng điều duy nhất họ không làm được là đưa bóng vào lưới thủ môn Marko Dmitrovic.\n\n  ![Sergio Gonzalez mừng bàn trong trận Barca 0-1 Leganes tối 15/12. Ảnh: EFE](https://vcdn1-thethao.vnecdn.net/2024/12/16/lega-1734305386-5441-1734305481.jpg?w=680&h=0&q=100&dpr=1&fit=crop&s=t8bFBKpj56rFSjWb46UHKw)\n\nSergio Gonzalez mừng bàn trong trận Barca 0-1 Leganes tối 15/12. Ảnh: _EFE_\n\nPhút thứ 10, Robert Lewandowski đệm bóng bật thủ môn từ cự ly 5 mét. Chừng 13 phút sau, Raphinha vô lê cận thành, nhưng Dmitrovic đẩy bóng vào xà ngang bật ra. Và đến phút 34, thủ môn từng dự World Cup 2018 và 2022 cùng tuyển Serbia lập hat-trick cản phá, khi dùng chân đẩy pha dứt điểm đối mặt của Lewandowski.\n\nDmitrovic vừa gia nhập Leganes hồi hè theo dạng chuyển nhượng tự do. Trước đó, thủ môn 32 tuổi từng giành Europa League 2022-2023 với Sevilla.\n\nLamine Yamal cũng chơi kém ấn tượng trong ngày khoe Cậu Bé Vàng, danh hiệu dành cho cầu thủ U21 hay nhất năm 2024. Cuối hiệp một, tiền đạo 17 tuổi thoát xuống rìa phải vòng cấm, nhưng sút vọt xà thay vì tạt cho đồng đội đệm bóng.\n\n  ![Raphinha tiếc sau pha dứt điểm trong trận Barca 0-1 Leganes tối 15/12. Ảnh: EFE](https://vcdn1-thethao.vnecdn.net/2024/12/16/ra-1734305397-4237-1734305481.jpg?w=680&h=0&q=100&dpr=1&fit=crop&s=sGnoNK4laK3QT3xE1zz-bA)\n\nRaphinha tiếc sau pha dứt điểm trong trận Barca 0-1 Leganes tối 15/12. Ảnh: _EFE_\n\nBarca tiếp tục chơi kém hiệu quả sau giờ nghỉ. Lewandowski, Yamal và Dani Olmo thiếu gắn kết trước khung thành Leganes, không tạo được thêm cơ hội nào trước khi bị thay ra. Đội chủ nhà cũng thất bại trong việc trông cậy vào các pha băng lên hỗ trợ tấn công của hậu vệ và tiền vệ trung tâm. Jules Kounde sút hai lần chệch khung thành, còn Fermin Lopez vô lê bật hậu vệ Leganes.\n\nTrong khi đó, đội khách chủ động lùi sâu tạo đám đông bảo vệ khung thành. Họ thành công trong việc duy trì kỷ luật phòng ngự và hạn chế không gian đối với [Barca](https://vnexpress.net/chu-de/barcelona-68).\n\nBarca tiếp tục dẫn đầu [La Liga](https://vnexpress.net/chu-de/la-liga-41) sau trận thua bất ngờ với 38 điểm sau 18 trận, nhưng bằng điểm và chơi nhiều hơn đội xếp thứ hai Atletico một trận. Real cũng mới chơi 17 trận và chỉ kém một điểm. Ngày 21/12, Barca sẽ tiếp Atletico, còn Real gặp Sevilla.\n\nLeganes leo từ 17 lên vị trí 15, nhờ ba điểm quý giá trước đội dẫn đầu. Đội bóng mới lên hạng hiện có 18 điểm sau 17 trận, hơn nhóm có nguy cơ xuống hạng bốn điểm.\n\n**Thanh Quý**\n\nVnExpress Newsletters\n\n### Đừng bỏ lỡ tin tức quan trọng!\n\nNhận tóm tắt tin tức nổi bật, hấp dẫn nhất 24 giờ qua trên VnExpress.")
	assert.Nil(t, err)

	assert.NotEmpty(t, ans)

	print(ans)
}
