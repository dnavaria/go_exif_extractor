package main

import (
	"fmt"
	CLI "go_image_exif_do/exif_cli"
	"os"
)

func main() {
	args := os.Args
	cliApp := CLI.NewCLI()
	result := cliApp.Run(args)
	if result != nil {
		fmt.Println("Go Image EXIF Extractor :: ", result)
	}

}
