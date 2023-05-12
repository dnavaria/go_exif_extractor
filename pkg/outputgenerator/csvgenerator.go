package outputgenerator

import (
	"encoding/csv"
	"fmt"
	"go_image_exif_do/pkg/exifextractor"
	"os"
	"path/filepath"
)

// CSVGenerator represents the CSV output generation functionality
type CSVGenerator struct {
}

// NewCSVGenerator creates a new instance of CSVGenerator
func NewCSVGenerator() *CSVGenerator {
	return &CSVGenerator{}
}

// Generate generates the output in CSV format
func (c *CSVGenerator) Generate(exifData []*exifextractor.ExifData, outputPath string) (string, error) {
	// Setting default output path
	if outputPath == "" {
		outputPath = "output.csv"
	} else {
		// If the output path ends with a slash, remove it
		if outputPath[len(outputPath)-1:] == "/" {
			outputPath = outputPath[:len(outputPath)-1]
			// Join the output path with the default output file name
			outputPath = filepath.Join(outputPath, "output.csv")
		} else {
			outputPath = filepath.Join(outputPath, "output.csv")
		}
	}
	// Create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return outputPath, fmt.Errorf("error :: failed to create CSV file: %v", err)
	}
	defer file.Close()

	// Create a new CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Image File Name", "Latitude", "Longitude"}
	err = writer.Write(header)
	if err != nil {
		return outputPath, fmt.Errorf("error :: failed to write header to CSV file: %v", err)
	}

	// Initialize metadata variables
	var imageFilePath string = ""
	var latitude string = "N/A"
	var longitude string = "N/A"

	// iterate over the EXIF data and write it to the CSV file
	for _, data := range exifData {
		// If the image file path is empty, set it to "exif data not found"
		imageFilePath = data.ImageFilePath
		latitude = "exif data not found"
		longitude = "exif data not found"

		// If the GPS data is not nil, set the latitude and longitude variables
		if data.GPS != nil {
			latitude = fmt.Sprintf("%v", data.GPS.Latitude)
			longitude = fmt.Sprintf("%v", data.GPS.Longitude)
		}
		// Write the record to the CSV file
		record := []string{
			imageFilePath,
			latitude,
			longitude,
		}
		err = writer.Write(record)
		if err != nil {
			fmt.Printf("error :: failed to write record to CSV file: %v", err)
		}
	}
	return outputPath, nil
}
