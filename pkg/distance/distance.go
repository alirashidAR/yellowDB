package distance

import (
	"math"

	"github.com/alirashidAR/yellowDB/pkg/vector"
)

type DistanceFunc func(a, b vector.Vector) float64

func Euclidean(a, b vector.Vector) float64 {
	sum := 0.0
	for i := range a.Values {
		diff := a.Values[i] - b.Values[i]
		sum += diff * diff
	}
	return math.Sqrt(sum)
}

func CosineSimilarity(a, b vector.Vector) float64 {
	dotProduct := 0.0
	magnitudeA := 0.0
	magnitudeB := 0.0

	for i := range a.Values {
		dotProduct += a.Values[i] * b.Values[i]
		magnitudeA += a.Values[i] * a.Values[i]
		magnitudeB += b.Values[i] * b.Values[i]
	}

	if magnitudeA == 0 || magnitudeB == 0 {
		return 0
	}

	return dotProduct / (math.Sqrt(magnitudeA) * math.Sqrt(magnitudeB))
}
