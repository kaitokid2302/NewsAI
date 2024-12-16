package ai

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/kaitokid2302/NewsAI/internal/config"
)

// summary text

type AIInfrast interface {
	Summarize(text string) (string, error)
}

type aiInfrast struct {
	provider config.Provider
}

func NewAIService(provider config.Provider) AIInfrast {
	return &aiInfrast{provider: provider}
}

func (g *aiInfrast) Summarize(text string) (string, error) {
	bodyMap := map[string]interface{}{
		"model": g.provider.Name,
		"messages": []map[string]interface{}{ // messages là một mảng
			{
				"role":    "user",
				"content": fmt.Sprintf(config.Global.Prompt, text),
			},
		},
	}

	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	er := encoder.Encode(bodyMap)
	if er != nil {
		return "", er
	}
	req, er := http.NewRequest("POST", "https://openrouter.ai/api/v1/chat/completions", &buf)
	if er != nil {
		return "", er
	}
	req.Header.Set("Authorization", "Bearer "+g.provider.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	res, er := client.Do(req)
	if er != nil {
		return "", er
	}
	defer res.Body.Close()
	var result map[string]interface{}
	er = json.NewDecoder(res.Body).Decode(&result)
	if er != nil {
		return "", er
	}
	choices, ok := result["choices"].([]interface{})
	if !ok {
		return "", errors.New("calling ai service failed")
	}
	if len(choices) == 0 {
		return "", errors.New("calling ai service failed")
	}
	messages, ok := choices[0].(map[string]interface{})["message"]
	if !ok {
		return "", errors.New("calling ai service failed")
	}
	content, ok := messages.(map[string]interface{})["content"]
	if !ok {
		return "", errors.New("calling ai service failed")
	}
	ans, ok := content.(string)
	if !ok {
		return "", errors.New("calling ai service failed")
	}
	return ans, nil
}
