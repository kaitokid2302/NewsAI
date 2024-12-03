package s3

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/kaitokid2302/NewsAI/internal/config"
	"math/rand/v2"
	"mime/multipart"
)

type UploadFileS3Service interface {
	UploadFile(file multipart.File) (string, error)
}

type UploadFileS3ServiceImpl struct {
	session *session.Session
}

func NewUploadFileS3Service(session *session.Session) *UploadFileS3ServiceImpl {
	return &UploadFileS3ServiceImpl{session: session}
}

func (u *UploadFileS3ServiceImpl) UploadFile(file multipart.File) (string, error) {
	fileName := "img" + fmt.Sprintf("%v", rand.Int32N(1000000000))
	uploader := s3manager.NewUploader(u.session)
	_, er := uploader.Upload(&s3manager.UploadInput{
		Bucket: &config.Global.Bucket,
		Key:    &fileName,
		Body:   file,
	})
	if er != nil {
		return "", er
	}
	link := fmt.Sprintf("https://%s.s3-%s.amazonaws.com/%s", config.Global.Bucket, u.session.Config.Region, fileName)
	return link, nil
}
