package user

import (
	"github.com/kaitokid2302/NewsAI/internal/database/model"
)

type UserReponse struct {
	StatusCode int        `json:"statusCode,omitempty"`
	Data       model.User `json:"data,omitempty"`
	Message    string     `json:"message,omitempty"`
}
