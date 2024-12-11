package topic

import (
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/kaitokid2302/NewsAI/internal/infrastructure/database"
	"github.com/stretchr/testify/assert"
)

func TestNewTopicRepository(t *testing.T) {
	// user1 to topic8

	config.InitAll()

	topicRepo := NewTopicRepository(database.InitDatabase())
	er := topicRepo.Subscribe(1, 8)
	assert.Nil(t, er)
	// subscribe to topic9

	er = topicRepo.Subscribe(1, 9)

	// all topic

	topics, er := topicRepo.AllTopicOfUser(1)
	assert.Nil(t, er)

	assert.Equal(t, 2, len(*topics))
	assert.Equal(t, uint(8), (*topics)[0].ID)
	assert.Equal(t, uint(9), (*topics)[1].ID)

	// Pháp luật
	// Giáo dục

	assert.Equal(t, "Pháp luật", (*topics)[0].Name)
	assert.Equal(t, "Giáo dục", (*topics)[1].Name)

	// unsubscribe topic9
	er = topicRepo.Unsubscribe(1, 9)
	assert.Nil(t, er)

	// all topic

	topics, er = topicRepo.AllTopicOfUser(1)
	assert.Nil(t, er)

	assert.Equal(t, 1, len(*topics))
	assert.Equal(t, uint(8), (*topics)[0].ID)
	assert.Equal(t, "Pháp luật", (*topics)[0].Name)

	// unsubscribe topic8
	er = topicRepo.Unsubscribe(1, 8)
	assert.Nil(t, er)

	// all topic

	topics, er = topicRepo.AllTopicOfUser(1)
	assert.Nil(t, er)

	assert.Equal(t, 0, len(*topics))
}

func TestFindTopicByName(t *testing.T) {
	config.InitAll()
	topicRepo := NewTopicRepository(database.InitDatabase())

	topic, er := topicRepo.FindTopicByName("Pháp luật")
	assert.Nil(t, er)
	assert.Equal(t, uint(8), topic.ID)
	assert.Equal(t, "Pháp luật", topic.Name)

	topic, er = topicRepo.FindTopicByName("Giáo dục")
	assert.Nil(t, er)
	assert.Equal(t, uint(9), topic.ID)
	assert.Equal(t, "Giáo dục", topic.Name)

	topic, er = topicRepo.FindTopicByName("Kinh tế")
	assert.NotNil(t, er)
	assert.Nil(t, topic)
}
