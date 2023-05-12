package tests

import (
	"go_image_exif_do/pkg/fshandler/directory"
	"os"
	"path/filepath"
	"testing"
)

func TestDirectory_GetFiles(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)

	// Create some test files
	createTestFile(t, filepath.Join(tempDir, "file1.txt"))
	createTestFile(t, filepath.Join(tempDir, "file2.jpg"))
	createTestFile(t, filepath.Join(tempDir, "file3.png"))
	createTestFile(t, filepath.Join(tempDir, "subdir/file4.txt"))

	dir := directory.NewDirectory()
	files, err := dir.GetFiles(tempDir)

	if err != nil {
		t.Errorf("Error getting files: %v", err)
	}

	expectedFiles := []string{
		filepath.Join(tempDir, "file2.jpg"),
		filepath.Join(tempDir, "file3.png"),
	}

	if len(files) != len(expectedFiles) {
		t.Errorf("Expected %d files, but got %d", len(expectedFiles), len(files))
	}

	for _, file := range expectedFiles {
		if !contains(files, file) {
			t.Errorf("Expected file %s not found", file)
		}
	}
}

func TestDirectory_IsDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)

	dir := directory.NewDirectory()

	// Test with an existing directory
	isDir, err := dir.IsDirectory(tempDir)
	if err != nil {
		t.Errorf("Error checking directory: %v", err)
	}
	if !isDir {
		t.Errorf("Expected %s to be a directory", tempDir)
	}

	// Test with a non-existing directory
	isDir, err = dir.IsDirectory(filepath.Join(tempDir, "nonexistent"))
	if err != nil {
		t.Errorf("Error checking directory: %v", err)
	}
	if isDir {
		t.Error("Expected nonexistent directory to be false")
	}
}

func TestDirectory_CreateDirectory(t *testing.T) {
	// Create a temporary directory for testing
	tempDir := createTempDir(t)
	defer os.RemoveAll(tempDir)

	dir := directory.NewDirectory()

	// Test creating a new directory
	newDir := filepath.Join(tempDir, "newdir")
	err := dir.CreateDirectory(newDir)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}

	// Check if the directory was created
	isDir, err := dir.IsDirectory(newDir)
	if err != nil {
		t.Errorf("Error checking directory: %v", err)
	}
	if !isDir {
		t.Errorf("Expected %s to be a directory", newDir)
	}

	// Test creating an existing directory
	err = dir.CreateDirectory(tempDir)
	if err != nil {
		t.Errorf("Error creating existing directory: %v", err)
	}

	// Check if the existing directory is still a directory
	isDir, err = dir.IsDirectory(tempDir)
	if err != nil {
		t.Errorf("Error checking existing directory: %v", err)
	}
	if !isDir {
		t.Errorf("Expected %s to be a directory", tempDir)
	}
}

// Helper function to create a temporary directory for testing
func createTempDir(t *testing.T) string {
	tempDir, err := os.MkdirTemp("", "directory_test")
	if err != nil {
		t.Fatalf("Error creating temporary directory: %v", err)
	}

	subdir := filepath.Join(tempDir, "subdir")
	err = os.MkdirAll(subdir, os.ModePerm)
	if err != nil {
		t.Fatalf("Error creating subdirectory: %v", err)
	}

	return tempDir
}

// Helper function to create a test file in the specified directory
func createTestFile(t *testing.T, path string) {
	file, err := os.Create(path)
	if err != nil {
		t.Fatalf("Error creating test file: %v", err)
	}
	defer file.Close()
}

// Helper function to check if a slice contains a specific string
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
