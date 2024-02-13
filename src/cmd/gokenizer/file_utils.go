package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func validateExtension(filename string, extension string) error {
	if !strings.HasSuffix(filename, extension) {
		return fmt.Errorf("error: file extension must be %s", extension)
	}
	return nil
}

func validateFileExists(filename string) error {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("error: file %s does not exist", filename)
	}
	return nil
}

func readFileAsString(filepath string) (string, error) {
	err := validateFileExists(filepath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = validateExtension(filepath, ".lua")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return "", fmt.Errorf("error reading file: %w", err)
	}

	return string(content), nil
}
