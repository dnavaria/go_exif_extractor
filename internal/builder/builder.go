package builder

import "fmt"

// Builder for command line utility
type UtilityBuilder struct {
	name         string
	inputPath    string
	outputFormat string
	outputPath   string
}

// Creates a new instance for UtilityBuilder
func NewUtilityBuilder() *UtilityBuilder {
	return &UtilityBuilder{}
}

// Set name of the utility
func (b *UtilityBuilder) SetName(name string) *UtilityBuilder {
	b.name = name
	return b
}

// Set input path  of the utility
func (b *UtilityBuilder) SetInputPath(inputPath string) *UtilityBuilder {
	b.inputPath = inputPath
	return b
}

// Set output format of the utility
func (b *UtilityBuilder) SetOutputFormat(outputFormat string) *UtilityBuilder {
	b.outputFormat = outputFormat
	return b
}

// Set output path of the utility
func (b *UtilityBuilder) SetOutputPath(outputPath string) *UtilityBuilder {
	b.outputPath = outputPath
	return b
}

type Utility struct {
	name         string
	inputPath    string
	outputFormat string
	outputPath   string
}

// Build cosntructs and returns the configured utility instance
func (b *UtilityBuilder) Build() (*Utility, error) {
	fmt.Println("Go Image EXIF DO")
	fmt.Printf("Name: %v\n", b.name)
	fmt.Printf("Input Path: %v\n", b.inputPath)
	fmt.Printf("Output Format: %v\n", b.outputFormat)
	fmt.Printf("Output Path: %v\n", b.outputPath)

	return &Utility{
		name:         b.name,
		inputPath:    b.inputPath,
		outputFormat: b.outputFormat,
		outputPath:   b.outputPath,
	}, nil
}

// Get Input Files
func (u *Utility) GetInputPath() string {
	return u.inputPath
}

// Get Output Format
func (u *Utility) GetOutputFormat() string {
	return u.outputFormat
}

// Get Output Path
func (u *Utility) GetOutputPath() string {
	return u.outputPath
}
