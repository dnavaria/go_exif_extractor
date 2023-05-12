package tests

import (
	"go_image_exif_do/pkg/fshandler/file"
	"os"
	"path/filepath"
	"testing"
)

// createMockTestFile creates a mock test file with the specified content
func createMockTestFile(t *testing.T, fp, content string) (string, func()) {
	tempFile := fp
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Error creating mock test file: %v", err)
	}

	// Return the path to the mock file and a cleanup function
	return tempFile, func() {
		os.Remove(tempFile)
	}
}

func TestFile_GetFilePath(t *testing.T) {
	f := file.NewFile().SetFilePath("/path/to/file.txt")
	expectedPath := "/path/to/file.txt"
	actualPath := f.GetFilePath()

	if actualPath != expectedPath {
		t.Errorf("Expected file path %s, but got %s", expectedPath, actualPath)
	}
}

func TestFile_GetFileExtension(t *testing.T) {
	f := file.NewFile().SetFilePath("/path/to/file.txt")
	expectedExtension := ".txt"
	actualExtension := f.GetFileExtension()

	if actualExtension != expectedExtension {
		t.Errorf("Expected file extension %s, but got %s", expectedExtension, actualExtension)
	}
}

func TestFile_OpenReadClose(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "mock_file.txt"
	filePath, cleanup := createMockTestFile(t, filepath.Join(tempDir, fileName), "Hello World")
	f := file.NewFile().SetFilePath(filePath)

	file, err := f.Open()
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	data := make([]byte, 10)
	_, err = file.Read(data)
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	err = file.Close()
	if err != nil {
		t.Errorf("Error closing file: %v", err)
	}
	cleanup()
}

func TestFile_Stat(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "mock_file.txt"
	filePath, cleanup := createMockTestFile(t, filepath.Join(tempDir, fileName), "Hello World")
	f := file.NewFile().SetFilePath(filePath)

	file, err := f.Open()
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}

	fileInfo, err := file.Stat()
	if err != nil {
		t.Errorf("Error getting file info: %v", err)
	}

	expectedName := "mock_file.txt"
	actualName := fileInfo.Name()
	if actualName != expectedName {
		t.Errorf("Expected file name %s, but got %s", expectedName, actualName)
	}
	f.Close()
	cleanup()
}

func TestFile_IsImageFile(t *testing.T) {
	f := file.NewFile()

	imageFile := "image.jpg"
	nonImageFile := "document.pdf"

	isImage := f.IsImageFile(imageFile)
	if !isImage {
		t.Errorf("Expected %s to be an image file", imageFile)
	}

	isImage = f.IsImageFile(nonImageFile)
	if isImage {
		t.Errorf("Expected %s to not be an image file", nonImageFile)
	}
}

func TestFile_GetFileMetadata(t *testing.T) {
	tempDir := t.TempDir()
	fileName := "mock_file.txt"
	filePath, cleanup := createMockTestFile(t, filepath.Join(tempDir, fileName), "Hello World")
	f := file.NewFile().SetFilePath(filePath)

	fileInfo, err := f.GetFileMetadata(filePath)
	if err != nil {
		t.Errorf("Error getting file metadata: %v", err)
	}

	expectedName := "mock_file.txt"
	actualName := fileInfo.Name()
	if actualName != expectedName {
		t.Errorf("Expected file name %s, but got %s", expectedName, actualName)
	}
	cleanup()
}
