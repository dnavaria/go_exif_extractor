package directory

import (
	"fmt"
	File "go_image_exif_do/pkg/fshandler/file"
	"os"
	"path/filepath"
)

// Directory for directory related operations
type Directory struct {
	filehandler *File.File
}

// NewDirectory creates a new instance of Directory
func NewDirectory() *Directory {
	return &Directory{
		filehandler: File.NewFile(),
	}
}

// GetFiles returns the list of files in the given directory
func (d *Directory) GetFiles(dir string) ([]string, error) {
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("error :: while traversing through directory %v\n", err)
			return err
		}
		if !info.IsDir() && d.filehandler.IsImageFile(path) {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error :: failed to get image files: %w", err)
	}
	return files, nil
}

// IsDirectory checks if the given path is a directory
func (d *Directory) IsDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}
	return fileInfo.IsDir(), nil
}

// CreateDirectory creates a directory at the given path
func (d *Directory) CreateDirectory(path string) error {
	// Remove extension if any
	ext := filepath.Ext(path)
	if ext != "" {
		path = path[:len(path)-len(ext)]
	}
	// Check if directory already exists
	isDirectory, err := d.IsDirectory(path)
	if err != nil {
		return err
	}
	if isDirectory {
		fmt.Println("Output directory found")
		return nil
	}
	fmt.Println("Creating output directory")
	// Create directory
	// check for / in the path
	if path[len(path)-1:] != "/" {
		path += "/"
	}
	dir := filepath.Dir(path)
	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
