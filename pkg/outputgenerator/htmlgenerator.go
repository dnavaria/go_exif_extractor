package outputgenerator

import (
	"fmt"
	"go_image_exif_do/pkg/exifextractor"
	"html/template"
	"os"
	"path/filepath"
)

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>EXIF Data Extraction</title>
</head>
<body>
    <h1>EXIF Data</h1>
    <table>
        <tr>
            <th>Image File</th>
            <th>Latitude</th>
            <th>Longitude</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.ImageFilePath}}</td>
            {{if .GPS}}
                <td>{{.GPS.Latitude}}</td>
                <td>{{.GPS.Longitude}}</td>
            {{else}}
                <td>exif data not found</td>
                <td>exif data not found</td>
            {{end}}
        </tr>
        {{end}}
    </table>
</body>
</html>
`

// HTMLGenerator represents the HTML output generation functionality.
type HTMLGenerator struct {
}

// NewHTMLGenerator creates a new instance of HTMLGenerator.
func NewHTMLGenerator() *HTMLGenerator {
	return &HTMLGenerator{}
}

// Generate generates an HTML file with the extracted EXIF data.
func (hg *HTMLGenerator) Generate(exifData []*exifextractor.ExifData, outputPath string) (string, error) {
	// setting default output path
	if outputPath == "" {
		outputPath = "output.html"
	} else {
		// if the output path ends with a slash, remove it
		if outputPath[len(outputPath)-1:] == "/" {
			outputPath = outputPath[:len(outputPath)-1]
			// join the output path with the default output file name
			outputPath = filepath.Join(outputPath, "output.html")
		} else {
			outputPath = filepath.Join(outputPath, "output.html")
		}
	}
	// create the output file
	file, err := os.Create(outputPath)
	if err != nil {
		return outputPath, fmt.Errorf("failed to create output file: %w", err)
	}
	defer file.Close()

	// parse the HTML template
	tmpl := template.Must(template.New("exif").Parse(htmlTemplate))
	// execute the template
	err = tmpl.Execute(file, exifData)
	if err != nil {
		return outputPath, fmt.Errorf("failed to execute HTML template: %w", err)
	}
	return outputPath, nil
}
