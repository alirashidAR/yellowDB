package storage

import (
	"encoding/gob"
	"errors"
	"os"
	"sync"

	"github.com/alirashidAR/yellowDB/pkg/vector"
)

type fileStorage struct {
	filePath string
	mutex    sync.RWMutex
	vectors  map[int]vector.Vector
}

// NewFileStorage creates a new instance of fileStorage with the specified file path.
func NewFileStorage(filePath string) *fileStorage {
	return &fileStorage{
		filePath: filePath,
		vectors:  make(map[int]vector.Vector),
	}
}

// Store saves a vector to the file and updates the in-memory map.
func (fs *fileStorage) Store(v vector.Vector) error {
	fs.mutex.Lock()
	defer fs.mutex.Unlock()

	// Store in the in-memory map
	fs.vectors[v.ID] = v

	// Append the vector to the file
	file, err := os.OpenFile(fs.filePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	return encoder.Encode(v)
}

// Retrieve fetches a vector by ID from the in-memory map.
func (fs *fileStorage) Retrieve(id int) (vector.Vector, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	v, exists := fs.vectors[id]
	if !exists {
		return vector.Vector{}, errors.New("vector not found")
	}
	return v, nil
}

// RetrieveAll loads all vectors from the file and updates the in-memory map.
func (fs *fileStorage) RetrieveAll() ([]vector.Vector, error) {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()

	file, err := os.Open(fs.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var vectors []vector.Vector
	decoder := gob.NewDecoder(file)
	for {
		var vec vector.Vector
		if err = decoder.Decode(&vec); err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		vectors = append(vectors, vec)
		fs.vectors[vec.ID] = vec // Update in-memory map
	}

	return vectors, nil
}

// Count returns the number of vectors stored in memory.
func (fs *fileStorage) Count() int {
	fs.mutex.RLock()
	defer fs.mutex.RUnlock()
	return len(fs.vectors)
}
