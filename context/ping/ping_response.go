package ping

type pingResponse struct {
	Pong string `json:"ping"`
}

func (this pingResponse) GetPong() string {
	return this.Pong
}
