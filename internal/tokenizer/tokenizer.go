package tokenizer

import (
	"strings"
)

type Tokenizer struct{}

// NewTokenizer creates a new instance of Tokenizer
func NewTokenizer() *Tokenizer {
	return &Tokenizer{}
}

// Tokenize splits the input text based on any whitespace character
func (t *Tokenizer) Tokenize(text string) []string {
	return strings.Fields(text)
}

// Embed creates a vector of float64 values based on the length of each token
func Embed(tokens []string) []float64 {
	vector := make([]float64, len(tokens))
	for i := range tokens {
		vector[i] = float64(len(tokens[i]))
	}
	return vector
}
