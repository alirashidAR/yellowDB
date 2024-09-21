package embedding

import (
	"math/rand"
	"strings"
	"time"
)


type Embedder struct{}


func NewEmbedder() *Embedder {
	rand.Seed(time.Now().UnixNano())
	return &Embedder{}
}


func (e *Embedder) Embed(sentence string) []float64 {
	words := strings.Fields(sentence)
	vector := make([]float64, 10)
	for i := range vector {
		vector[i] = float64(len(words)) * rand.Float64() 
	}
	return vector
}
