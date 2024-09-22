package index

import "github.com/alirashidAR/yellowDB/pkg/vector"

type VectorIndex interface {
	Build(vectors []vector.Vector) error
	Search(query vector.Vector, k int) ([]vector.Vector, error)
}
