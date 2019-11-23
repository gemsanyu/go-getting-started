package user

type getTokenRequest struct {
	Username    string `json:"username"`
	Password string `json:"password"`
}

func (this getTokenRequest) GetUsername() string {
	return this.Username
}

func (this getTokenRequest) GetPassword() string {
	return this.Password
}
