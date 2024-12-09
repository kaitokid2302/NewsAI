package ai

import (
	"testing"

	"github.com/kaitokid2302/NewsAI/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestAIService(t *testing.T) {
	config.InitAll()
	provider := config.Global.AI.Provider[0]

	aiService := NewAIService(provider)

	ans, err := aiService.Summarize("Hello")
	assert.Nil(t, err)

	assert.NotEmpty(t, ans)

	print(ans)
}
