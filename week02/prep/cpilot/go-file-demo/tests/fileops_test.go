package tests

import (
	"os"
	"testing"

	"your_module_name/src" // replace with your actual module name
)

func TestCreateFile(t *testing.T) {
	filename := "testfile.txt"
	err := src.CreateFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = os.Stat(filename)
	if os.IsNotExist(err) {
		t.Fatalf("Expected file %s to exist, got %v", filename, err)
	}

	// Clean up
	os.Remove(filename)
}

func TestWriteFile(t *testing.T) {
	filename := "testfile.txt"
	content := []byte("Hello, World!")

	err := src.WriteFile(filename, content)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if string(data) != string(content) {
		t.Fatalf("Expected content %s, got %s", content, data)
	}

	// Clean up
	os.Remove(filename)
}

func TestReadFile(t *testing.T) {
	filename := "testfile.txt"
	content := []byte("Hello, World!")

	// First, create and write to the file
	err := src.WriteFile(filename, content)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	readContent, err := src.ReadFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if string(readContent) != string(content) {
		t.Fatalf("Expected content %s, got %s", content, readContent)
	}

	// Clean up
	os.Remove(filename)
}

func TestDeleteFile(t *testing.T) {
	filename := "testfile.txt"

	// Create the file first
	err := src.CreateFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	err = src.DeleteFile(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		t.Fatalf("Expected file %s to be deleted, got %v", filename, err)
	}
}