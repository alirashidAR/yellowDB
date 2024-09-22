package index

import (
	"sort"

	"github.com/alirashidAR/yellowDB/pkg/distance"
	"github.com/alirashidAR/yellowDB/pkg/vector"
)

type LinearIndex struct {
	vectors  []vector.Vector
	distFunc distance.DistanceFunc
}

func NewLinearIndex(distFunc distance.DistanceFunc) *LinearIndex {
	return &LinearIndex{distFunc: distFunc}
}

func (li *LinearIndex) Build(vectors []vector.Vector) error {
	li.vectors = vectors
	return nil
}

func (li *LinearIndex) Search(query vector.Vector, k int) ([]vector.Vector, error) {
	type result struct {
		vector vector.Vector
		dist   float64
	}

	results := make([]result, len(li.vectors))
	for i, vec := range li.vectors {
		results[i] = result{vector: vec, dist: li.distFunc(query, vec)}
	}

	sort.Slice(results, func(i, j int) bool {
		return results[i].dist < results[j].dist
	})

	topK := make([]vector.Vector, k)
	for i := 0; i < k; i++ {
		topK[i] = results[i].vector
	}

	return topK, nil
}
