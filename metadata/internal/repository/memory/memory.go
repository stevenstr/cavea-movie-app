package memory

import (
	"mymovie_applctn/metadata/pkg/model"
	"sync"
)

type Repository struct {
	sync.RWMutex
	data map[string]*model.MetaData
}

func New() *Repository {
	return &Repository{data: map[string]*model.MetaData{}}
}

func Get() {}
func Put() {}
