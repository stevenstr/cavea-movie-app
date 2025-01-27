package rating

import (
	"context"
	"errors"

	"movie.com/metadata/internal/repository"
	"movie.com/rating/pkg/model"
)

var ErrNotFound = errors.New("ratings not found for a record")

type ratingRepository interface {
	Get(ctx context.Context, recordID model.RecordID,
		recordType model.RecordType) ([]model.Rating, error)
	Put(ctx context.Context, recordID model.RecordID,
		recordType model.RecordType, rating *model.Rating) error
}

type Controller struct {
	repo ratingRepository
}

func New(repo ratingRepository) *Controller {
	return &Controller{repo: repo}
}

func (c *Controller) GetAggregatedRating(ctx context.Context, recordID model.RecordID,
	recordType model.RecordType) (float64, error) {
	ratings, err := c.repo.Get(ctx, recordID, recordType)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		return 0, ErrNotFound
	} else if err != nil {
		return 0, err
	}
	sum := float64(0)

	for _, v := range ratings {
		sum += float64(v.Value)
	}
	return sum / float64(len(ratings)), nil

}

func (c *Controller) PutRating(ctx context.Context, recordId model.RecordID,
	recordType model.RecordType, rating *model.Rating) error {

	return c.repo.Put(ctx, recordId, recordType, rating)

}
