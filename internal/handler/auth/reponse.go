package auth

import "github.com/kaitokid2302/NewsAI/internal/database"

type Response struct {
	StatusCode int           `json:"statusCode,omitempty"`
	Data       database.User `json:"data,omitempty"`
	Er         string        `json:"error,omitempty"`
	Message    string        `json:"message,omitempty"`
	JwtToken   string        `json:"jwt,omitempty"`
}
