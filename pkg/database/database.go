package database

import (
	"errors"

	"github.com/alirashidAR/yellowDB/internal/index"
	"github.com/alirashidAR/yellowDB/internal/storage"
	"github.com/alirashidAR/yellowDB/pkg/distance"
	"github.com/alirashidAR/yellowDB/pkg/vector"
)

type Config struct {
	UseStorage   bool
	StoragePath  string
	IndexType    string
	DistanceType string
}

type VectorDB struct {
	memStore   []vector.Vector
	storage    storage.VectorStorage
	index      index.VectorIndex
	useStorage bool
}

func New(config Config) (*VectorDB, error) {
	var distFunc distance.DistanceFunc
	switch config.DistanceType {
	case "euclidean":
		distFunc = distance.Euclidean
	case "cosine":
		distFunc = distance.CosineSimilarity
	default:
		return nil, errors.New("unsupported distance type")
	}

	idx := index.NewLinearIndex(distFunc)

	var stor storage.VectorStorage
	if config.UseStorage {
		stor = storage.NewFileStorage(config.StoragePath)
	}

	return &VectorDB{
		index:      idx,
		storage:    stor,
		useStorage: config.UseStorage,
	}, nil
}

func (vdb *VectorDB) Add(v vector.Vector) error {
	vdb.memStore = append(vdb.memStore, v)
	if vdb.useStorage {
		return vdb.storage.Store(v)
	}
	return nil
}

func (vdb *VectorDB) BuildIndex() error {
	return vdb.index.Build(vdb.memStore)
}

func (vdb *VectorDB) NearestNeighbors(query vector.Vector, k int) ([]vector.Vector, error) {
	return vdb.index.Search(query, k)
}
