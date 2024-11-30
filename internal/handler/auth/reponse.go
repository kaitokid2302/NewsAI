package auth

type Response struct {
	StatusCode int          `json:"statusCode,omitempty"`
	Data       ResponseData `json:"data,omitempty"`
	Er         string       `json:"error,omitempty"`
	Message    string       `json:"message,omitempty"`
	JwtToken   string       `json:"jwt,omitempty"`
}

type ResponseData struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
