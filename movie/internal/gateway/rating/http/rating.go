package http

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"movie.com/movie/internal/gateway"
	"movie.com/rating/pkg/model"
)

type Gateway struct {
	addr string
}

func New(address string) *Gateway {
	return &Gateway{addr: address}
}

func (g *Gateway) GetAggregatedRating(ctx context.Context, recordID model.RecordID,
	recordType model.RecordType) (float64, error) {
	req, err := http.NewRequest(http.MethodGet, g.addr+"/rating", nil)
	if err != nil {
		return 0, err
	}

	req = req.WithContext(ctx)

	values := req.URL.Query()

	values.Add("id", string(recordID))
	values.Add("type", string(recordType))

	req.URL.RawQuery = values.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return 0, gateway.ErrNotFound
	} else if resp.StatusCode/100 != 2 {
		return 0, fmt.Errorf("non 2xx response: %v", resp)
	}

	var v float64
	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return 0, err
	}

	return v, nil

}

func (g *Gateway) PutRating(ctx context.Context, recordId model.RecordID,
	recordType model.RecordType, rating *model.Rating) error {

}
