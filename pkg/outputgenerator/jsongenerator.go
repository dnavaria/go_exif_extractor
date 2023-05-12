package outputgenerator

import (
	"encoding/json"
	"fmt"
	"go_image_exif_do/pkg/exifextractor"
	"os"
	"path/filepath"
)

// JSONGenerator represents the JSON output generation functionality.
type JSONGenerator struct {
}

// NewJSONGenerator creates a new instance of JSONGenerator.
func NewJSONGenerator() *JSONGenerator {
	return &JSONGenerator{}
}

// Generate generates a JSON file with the extracted EXIF data.
func (jg *JSONGenerator) Generate(exifData []*exifextractor.ExifData, outputPath string) (string, error) {
	// setting default output path
	if outputPath == "" {
		outputPath = "output.json"
	} else {
		// if the output path ends with a slash, remove it
		if outputPath[len(outputPath)-1:] == "/" {
			outputPath = outputPath[:len(outputPath)-1]
			// join the output path with the default output file name
			outputPath = filepath.Join(outputPath, "output.json")
		} else {
			outputPath = filepath.Join(outputPath, "output.json")
		}
	}
	// create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return outputPath, fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// encode the EXIF data as JSON
	encoder := json.NewEncoder(file)

	// set indentation for the output JSON
	encoder.SetIndent("", "  ")

	// encode the EXIF data
	err = encoder.Encode(exifData)
	if err != nil {
		return outputPath, fmt.Errorf("failed to encode JSON: %w", err)
	}
	return outputPath, nil
}
