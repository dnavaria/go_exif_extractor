package exif_cli

import (
	"flag"
	"fmt"
	"go_image_exif_do/internal/builder"
	"go_image_exif_do/pkg/exifextractor"
	Directory "go_image_exif_do/pkg/fshandler/directory"
	"go_image_exif_do/pkg/outputgenerator"
	"os"
)

// CLI represents the command line interface.
type CLI struct {
	exifExtractor    *exifextractor.ExifExtractor
	outputGenerator  *outputgenerator.OutputGenerator
	directoryHandler *Directory.Directory

	// Command line flags
	inputPath    string
	outputPath   string
	outputFormat string
}

// NewCLI creates a new instance of CLI.
func NewCLI() *CLI {
	return &CLI{
		exifExtractor:    exifextractor.NewExifExtractor(),
		outputGenerator:  outputgenerator.NewOutputGenerator(),
		directoryHandler: Directory.NewDirectory(),
	}
}

// parseFlags parses the command line flags.
func (cli *CLI) parseFlags() {
	// Define and parse command line flags
	flag.StringVar(&cli.inputPath, "input_path", "", "Path to the input file or directory")
	flag.StringVar(&cli.outputPath, "output_path", "", "Path to the output directory")
	flag.StringVar(&cli.outputFormat, "output_format", "", "Output format (e.g., json, csv, html)")

	flag.Parse()

	// Check if required flags are provided
	if cli.inputPath == "" || cli.outputPath == "" || cli.outputFormat == "" {
		fmt.Println("Missing required flags. Please provide values for input_path, output_path, and output_format.")
		flag.Usage()
		os.Exit(1)
	}
}

// Run runs the CLI.
func (cli *CLI) Run(args []string) error {
	// Parse command line flags
	cli.parseFlags()

	// Use the utility builder to set the configuration
	utility, err := builder.NewUtilityBuilder().
		SetName("go_image_exif_do").
		SetInputPath(cli.inputPath).
		SetOutputPath(cli.outputPath).
		SetOutputFormat(cli.outputFormat).
		Build()

	if err != nil {
		return err
	}

	// Process input files/directory and extract EXIF data
	imageFiles, err := cli.directoryHandler.GetFiles(utility.GetInputPath())
	if err != nil {
		return err
	}

	// Extract EXIF data
	var exifData []*exifextractor.ExifData
	for _, imagePath := range imageFiles {
		gpsData, err := cli.exifExtractor.ExtractGPSData(imagePath)
		if err != nil {
			// If error occurs, add nil GPS data
			exifData = append(exifData, &exifextractor.ExifData{
				ImageFilePath: imagePath,
				GPS:           nil,
			})
			continue
		}
		exifData = append(exifData, &exifextractor.ExifData{
			ImageFilePath: imagePath,
			GPS:           gpsData,
		})
	}
	// Generate output
	outputPath := utility.GetOutputPath()
	// Create output directory if it does not exist
	cli.directoryHandler.CreateDirectory(outputPath)
	// Generate output based on the output format
	outputPath, err = cli.outputGenerator.GenerateOutput(utility.GetOutputFormat(), exifData, outputPath)
	if err != nil {
		return err
	}
	fmt.Println("Output generated :: ", outputPath, " :: success")
	return nil
}
