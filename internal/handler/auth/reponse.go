package auth

type RegisterResponse struct {
	StatusCode int
	Data       RegisterResponseData
	Er         string
	Message    string
}

type RegisterResponseData struct {
	Email string
	Name  string `json:"name"`
}
