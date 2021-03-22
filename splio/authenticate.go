package splio

type Authenticate struct {
	ApiKey string `json:"api_key"`
}

type AuthenticateResponse struct {
	Token string `json:"token"`
}