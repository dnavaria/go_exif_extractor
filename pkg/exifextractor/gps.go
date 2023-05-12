package exifextractor

import "fmt"

// GPS struct represents the GPS data
// It contains the latitude and longitude values
type GPS struct {
	Latitude  float64
	Longitude float64
}

// NewGPS creates a new instance of GPS
// It returns an error if the latitude or longitude values are invalid
// The latitude value must be between -90 and 90
// The longitude value must be between -180 and 180
func NewGPS(latitude, longitude float64) (*GPS, error) {
	// Validate latitude and longitude values
	if latitude < -90 || latitude > 90 {
		return nil, fmt.Errorf("error :: invalid latitude value: %f", latitude)
	}

	if longitude < -180 || longitude > 180 {
		return nil, fmt.Errorf("error :: invalid longitude value: %f", longitude)
	}

	return &GPS{
		Latitude:  latitude,
		Longitude: longitude,
	}, nil
}
