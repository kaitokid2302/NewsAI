package auth

type Response struct {
	StatusCode int
	Data       ResponseData `json:"data,omitempty"`
	Er         string
	Message    string
	JwtToken   string `json:"jwt,omitempty"`
}

type ResponseData struct {
	Email string `json:"email,omitempty"`
	Name  string `json:"name,omitempty"`
}
