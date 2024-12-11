package topic

import (
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/aws"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/kaitokid2302/NewsAI/internal/repository/topic"
	user2 "github.com/kaitokid2302/NewsAI/internal/repository/user"
	"github.com/kaitokid2302/NewsAI/internal/service/s3"
	"github.com/kaitokid2302/NewsAI/internal/service/user"
	"github.com/stretchr/testify/assert"
)

func TestTopicSubscribeAndUnsubscribe(t *testing.T) {

	config.InitAll()
	db := database.InitDatabase()
	topicService := NewTopicService(user.NewUserService(s3.NewUploadFileS3Service(aws.AwsInit()), user2.NewUserRepo(db)), topic.NewTopicRepository(db))

	er := topicService.Subscribe("truonglamthientai321@gmail.com", "Pháp luật")
	assert.Nil(t, er)

	er = topicService.Unsubscribe("truonglamthientai321@gmail.com", "Kinh tế")
	assert.NotNil(t, er)

	er = topicService.Subscribe("truonglamthientai321@gmail.com", "Du lịch")

	assert.Nil(t, er)

	allTopics, er := topicService.AllTopic("truonglamthientai321@gmail.com")
	assert.Nil(t, er)

	assert.Equal(t, 2, len(*allTopics))
	assert.Equal(t, "Pháp luật", (*allTopics)[0].Name)
	assert.Equal(t, "Du lịch", (*allTopics)[1].Name)

	er = topicService.Unsubscribe("truonglamthientai321@gmail.com", "Du lịch")
	assert.Nil(t, er)

	allTopics, er = topicService.AllTopic("truonglamthientai321@gmail.com")
	assert.Nil(t, er)

	assert.Equal(t, 1, len(*allTopics))

	assert.Equal(t, "Pháp luật", (*allTopics)[0].Name)

	er = topicService.Unsubscribe("truonglamthientai321@gmail.com", "Du lịch")

	assert.NotNil(t, er)

	er = topicService.Unsubscribe("truonglamthientai321@gmail.com", "Pháp luật")
	assert.Nil(t, er)

	allTopics, er = topicService.AllTopic("truonglamthientai321@gmail.com")
	assert.Nil(t, er)

	assert.Equal(t, 0, len(*allTopics))

}
