package file

import (
	"io/fs"
	"os"
	"path/filepath"
)

// File for file operations
type File struct {
	file_path string
	file      *os.File
}

// NewFile creates a new instance of File
func NewFile() *File {
	return &File{}
}

// SetFilePath sets the file path
func (f *File) SetFilePath(path string) *File {
	f.file_path = path
	return f
}

// GetFilePath returns the file path
func (f *File) GetFilePath() string {
	return f.file_path
}

// GetFileExtension returns the file extension
func (f *File) GetFileExtension() string {
	return filepath.Ext(f.file_path)
}

// Open opens the file for reading
func (f *File) Open() (fs.File, error) {
	file, err := os.Open(f.file_path)
	if err != nil {
		return nil, err
	}
	f.file = file
	return f, nil
}

// Close closes the file
func (f *File) Close() error {
	err := f.file.Close()
	f.file = nil
	return err
}

// Read reads from the file
func (f *File) Read(b []byte) (int, error) {
	return f.file.Read(b)
}

// Stat returns file information
func (f *File) Stat() (fs.FileInfo, error) {
	return f.file.Stat()
}

// isImageFile checks if the given file path has an image file extension.
func (f *File) IsImageFile(path string) bool {
	// using map for faster lookup
	extensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
	}
	ext := filepath.Ext(path)
	_, ok := extensions[ext]
	return ok
}

// GetFileMetadata returns the metadata of the given file
func (f *File) GetFileMetadata(path string) (fs.FileInfo, error) {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	return fileInfo, nil
}
