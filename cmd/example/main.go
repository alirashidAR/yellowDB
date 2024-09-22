// cmd/example/main.go
package main

import (
	"fmt"
	"log"

	"github.com/alirashidAR/yellowDB/pkg/database"
	"github.com/alirashidAR/yellowDB/internal/embedding" // Embedder package for generating embeddings
	"github.com/alirashidAR/yellowDB/pkg/vector"
	"github.com/alirashidAR/yellowDB/internal/tokenizer"
)

func main() {
	// Configuring the Vector Database with storage and cosine similarity
	config := database.Config{
		UseStorage:   true,               // Enable persistent storage
		StoragePath:  "vectors.db",       // File path for storing vectors
		IndexType:    "linear",           // Using a linear search index
		DistanceType: "cosine",           // Use cosine similarity for searching
	}

	// Create a new instance of the vector database
	db, err := database.New(config)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	// Define some sample sentences
	sentences := []string{
		"Hello world",
		"Goodbye world",
		"Golang is great",
		"Programming is fun",
		"Machine learning is awesome",
		"Data science is cool",
		"Python is popular",
	}

	embedder := embedding.NewEmbedder()
	tokenizer := tokenizer.NewTokenizer()
	val := tokenizer.Tokenize("Hello World my name is Ali")
	fmt.Println(val)
	fmt.Println(len(val))

	

	// Add embedded vectors to the database
	for i, sentence := range sentences {
		// Embed the sentence using the embedder package
		vectorValues := embedder.Embed(sentence)
		if err != nil {
			log.Fatalf("Failed to embed sentence: %v", err)
		}

		// Create a vector using the embedded values
		v := vector.Vector{ID: i + 1, Values: vectorValues, Text: sentence}

		// Add the vector to the database
		if err := db.Add(v); err != nil {
			log.Printf("Failed to add vector ID %d: %v", v.ID, err)
		}
	}

	// Build the index from the added vectors
	if err := db.BuildIndex(); err != nil {
		log.Fatalf("Failed to build index: %v", err)
	}

	// Define a query sentence to find similar vectors
	queryText := "Golang programming"

	// Embed the query sentence
	queryVectorValues := embedder.Embed(queryText)
	if err != nil {
		log.Fatalf("Failed to embed query sentence: %v", err)
	}

	// Create the query vector
	query := vector.Vector{Values: queryVectorValues, Text: queryText}

	// Find the 2 nearest neighbors to the query vector
	neighbors, err := db.NearestNeighbors(query, 2)
	if err != nil {
		log.Fatalf("Failed to find nearest neighbors: %v", err)
	}

	// Print the results
	fmt.Println("Nearest Neighbors:")
	for _, neighbor := range neighbors {
		fmt.Printf("ID: %d, Values: %v, Text: %s\n", neighbor.ID, neighbor.Values, neighbor.Text)
	}
}
