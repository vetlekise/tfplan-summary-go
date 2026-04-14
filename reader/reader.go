package reader

import (
	"log"
	"os"
	"path/filepath"
)

// Validates the file extension '.json' of the provided file and reads it.
func ReadPlan(data string) []byte {

	// Validate extension
	fileExtension := filepath.Ext(data)
	if fileExtension != ".json" {
		panic("File extension must be '.json'!")
	}

	// Open the JSON file
	file, err := os.ReadFile(data)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}

	return file
}
