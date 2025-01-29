package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"movie.com/metadata/pkg/model"
	"movie.com/movie/internal/gateway"
)

type Gateway struct {
	addr string
}

func New(address string) *Gateway {
	return &Gateway{addr: address}
}

func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/metadata", nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	values := req.URL.Query()
	values.Add("id", id)
	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return nil, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non 2xx response: %v", resp)
	}

	var v *model.Metadata
	if err := json.NewDecoder(req.Body).Decode(&v); err != nil {
		return nil, err
	}

	return v, nil
}
