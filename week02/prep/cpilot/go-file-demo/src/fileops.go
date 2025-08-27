package main

import (
    "os"
)

// CreateFile creates a new file with the specified name and content.
func CreateFile(filename string, content []byte) error {
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(content)
    return err
}

// ReadFile reads the content of the specified file.
func ReadFile(filename string) ([]byte, error) {
    content, err := os.ReadFile(filename)
    return content, err
}

// WriteFile writes content to the specified file.
func WriteFile(filename string, content []byte) error {
    file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
    if err != nil {
        return err
    }
    defer file.Close()

    _, err = file.Write(content)
    return err
}

// DeleteFile deletes the specified file.
func DeleteFile(filename string) error {
    return os.Remove(filename)
}