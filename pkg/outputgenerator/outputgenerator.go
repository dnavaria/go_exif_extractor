package outputgenerator

import (
	"fmt"
	"go_image_exif_do/pkg/exifextractor"
)

// OutputGenerator represents the output generation functionality.
// It is a wrapper around the different output generators.
// It is used to generate the desired output format based on the provided format and exifData.
type OutputGenerator struct {
	csvGenerator  *CSVGenerator
	jsonGenerator *JSONGenerator
	htmlGenerator *HTMLGenerator
}

// NewOutputGenerator creates a new instance of OutputGenerator.
func NewOutputGenerator() *OutputGenerator {
	return &OutputGenerator{
		csvGenerator:  NewCSVGenerator(),
		jsonGenerator: NewJSONGenerator(),
		htmlGenerator: NewHTMLGenerator(),
	}
}

// GenerateOutput generates the desired output format based on the provided format and exifData.
func (og *OutputGenerator) GenerateOutput(format string, exifData []*exifextractor.ExifData, outputPath string) (string, error) {
	// generate the output based on the provided format
	switch format {
	case "CSV", "csv":
		return og.csvGenerator.Generate(exifData, outputPath)
	case "JSON", "json":
		return og.jsonGenerator.Generate(exifData, outputPath)
	case "HTML", "html":
		return og.htmlGenerator.Generate(exifData, outputPath)
	default:
		fmt.Println("No valid output format provided, defaulting to CSV")
		return og.csvGenerator.Generate(exifData, outputPath)
	}
}
