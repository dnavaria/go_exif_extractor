# Image Exif Data Extraction Utility

## Description

This project is a command-line interface (CLI) tool for extracting EXIF data from images and generating output in various formats.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Makefile](#makefile)

## Installation

To install and run this project locally, follow these steps:

1. Change to the project directory:

   ```shell
   cd go_image_exif_do
   ```

2. Build the project using the Makefile:

   ```shell
   make build
   ```

## Usage

To use the project, execute the generated binary file with the appropriate command-line flags. The available flags are:

- `--input_path`: Path to the input file or directory.
- `--output_path`: Path to the output directory.
- `--output_format`: Output format (e.g., json, csv, html).

Example usage:

```shell
./go_image_exif_do --input_path /path/to/input --output_path /path/to/output --output_format csv
```

Replace `/path/to/input` with the actual path to your input file or directory. Similarly, replace `/path/to/output` with the desired path for the output directory. You can also change `csv` to your preferred output format, such as `json` or `html`.

## Makefile

This project includes a Makefile that provides several targets to help with building, cleaning, and testing the project. Here are the available targets:

- `all`: The default target that builds the project (depends on `build`).
- `build`: Compiles the project using `go build` and creates the binary.
- `clean`: Cleans up the project by running `go clean` and removing the binary.
- `test`: Runs tests using `go test` in the specified directory.

To execute a target, navigate to the project directory in your terminal and use the `make` command followed by the target name. For example, to build the project, run:

```shell
make build
```
