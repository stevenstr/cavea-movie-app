package memory

import (
	"context"
	"sync"

	"github.com/stevenstr/cavea-movie-app/metadata/internal/repository"
	"github.com/stevenstr/cavea-movie-app/metadata/pkg/model"
)

// Repository defines in memory reprisentation set of metadata.
type Repository struct {
	mx   sync.RWMutex
	data map[string]*model.Metadata
}

// New create new repository in memory.
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrive the metadata from repo by id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put add the new metadata by proven id.
func (r *Repository) Put(_ context.Context, id string, metadata *model.Metadata) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.data[id] = metadata

	return nil
}
