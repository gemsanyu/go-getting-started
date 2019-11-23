package ping

type PingUsecase interface {
	Ping(PingRequest) (PingResponse, error)
}

type PingRequest interface {
}

type PingResponse interface {
	GetPong() string
}

type pingUsecase struct{}

func NewPingUsecase() PingUsecase {
	return &pingUsecase{}
}

func (this *pingUsecase) Ping(req PingRequest) (PingResponse, error) {
	resp := pingResponse{"Pong"}
	return resp, nil
}
