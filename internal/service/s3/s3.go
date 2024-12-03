package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kaitokid2302/NewsAI/internal/config"
	"math/rand/v2"
	"mime/multipart"
	"net/http"
)

type UploadFileS3Service interface {
	UploadFile(name string, file multipart.File) (string, error)
}

type UploadFileS3ServiceImpl struct {
	session *session.Session
}

func NewUploadFileS3Service(session *session.Session) *UploadFileS3ServiceImpl {
	return &UploadFileS3ServiceImpl{session: session}
}

func (u *UploadFileS3ServiceImpl) UploadFile(name string, file multipart.File) (string, error) {

	buffer := make([]byte, 512)
	_, err := file.Read(buffer)
	if err != nil {
		return "", err
	}

	// Seek về đầu file sau khi đã đọc
	file.Seek(0, 0)

	// Detect content type
	contentType := http.DetectContentType(buffer)

	fileName := "img" + fmt.Sprintf("%v", rand.Int32N(1000000000)) + "-" + fmt.Sprintf("%v", rand.Int32N(1000000000))
	uploader := s3manager.NewUploader(u.session)
	_, er := uploader.Upload(&s3manager.UploadInput{
		Bucket:      &config.Global.Bucket,
		Key:         &fileName,
		Body:        file,
		ContentType: aws.String(contentType),
	},
	)
	if er != nil {
		return "", er
	}
	link := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", config.Global.Bucket, *u.session.Config.Region, fileName)
	return link, nil
}
