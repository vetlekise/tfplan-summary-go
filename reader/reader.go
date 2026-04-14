// Package reader implements functions for reading Terraform Plan files.
package reader

import (
	"log"
	"os"
	"path/filepath"
)

// ReadPlan validates that the given path has a .json extension and returns its contents.
// It panics if the extension is invalid and calls [log.Fatalf] if the file cannot be read.
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
