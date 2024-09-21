// internal/storage/storage.go
package storage

import "github.com/alirashidAR/yellowDB/pkg/vector"

// VectorStorage is an interface for storing vectors.
type VectorStorage interface {
	Store(v vector.Vector) error
	RetrieveAll() ([]vector.Vector, error)
	Count() int
}
