package http

type Gateway struct {
	addr string
}

func New(address string) *Gateway {
	return &Gateway{addr: address}
}
