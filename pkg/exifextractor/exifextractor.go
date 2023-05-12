package exifextractor

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

// ExifExtractor represents the EXIF data extractino functionality
// It is used to extract the GPS data from the given image file
// It is a wrapper around the goexif/exif library
type ExifExtractor struct {
}

// NewExifExtractor creates a new instance of ExifExtractor
func NewExifExtractor() *ExifExtractor {
	return &ExifExtractor{}
}

// ExtractGPSData extracts the GPS data from the given image file
// It returns the GPS data as a GPS struct
// It returns an error if the GPS data could not be extracted or decoded or if the image file could not be opened or read from disk
func (e *ExifExtractor) ExtractGPSData(path string) (*GPS, error) {
	// Open image file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error :: opening file: %v", err)
	}
	defer file.Close()

	// Decode EXIF data
	exifData, err := exif.Decode(file)
	if err != nil {
		return nil, fmt.Errorf("error :: decoding EXIF data: %v", err)
	}

	// Extract GPS data
	latitude, longitude, err := e.ExtractGPSFromExif(exifData)
	if err != nil {
		return nil, fmt.Errorf("error :: extracting GPS data: %v", err)
	}

	return NewGPS(latitude, longitude)
}

// ExtractGPSFromExif extracts the GPS data from the given EXIF data
// It returns the GPS data as latitude and longitude
// It returns an error if the GPS data could not be extracted or decoded
// It returns 0, 0 if the GPS data could not be found
func (ee *ExifExtractor) ExtractGPSFromExif(exifData *exif.Exif) (float64, float64, error) {
	latitude, longitude, err := exifData.LatLong()
	if err != nil {
		return 0, 0, nil
	}
	return latitude, longitude, nil
}
