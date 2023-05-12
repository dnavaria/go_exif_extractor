package exifextractor

// ExifData represents the extracted EXIF data
// It contains the image file path and the GPS data
// The GPS data is represented as a GPS struct
// The GPS data can be nil if the GPS data could not be extracted or decoded
// The image file path can be empty if the image file could not be opened or read from disk
type ExifData struct {
	ImageFilePath string
	GPS           *GPS
}
