package storage

import "github.com/alirashidAR/yellowDB/pkg/vector"

// VectorStorage defines the interface for vector storage systems.
type VectorStorage interface {
	Store(v vector.Vector) error
	Retrieve(id int) (vector.Vector, error)
	RetrieveAll() ([]vector.Vector, error)
	Count() int
}
