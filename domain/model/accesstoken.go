package model

type AccessToken struct {
	AccessToken string      `json:"access_token"`
	TokenType   string      `json:"type_type"`
	ExpiresIn   int64       `json:"expires_in"`
	Role        string      `json:"role"`
	Data        interface{} `json:"data"`
}
