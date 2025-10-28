package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// expandTilde expands a path starting with ~ to the home directory
func expandTilde(dirPath string) (string, error) {
	if len(dirPath) > 0 && dirPath[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		if len(dirPath) > 1 {
			return filepath.Join(home, dirPath[1:]), nil
		}
		return home, nil
	}
	return dirPath, nil
}

// resolvePath resolves a directory path to an absolute path, expanding tilde if needed
func resolvePath(dirPath string) (string, error) {
	expandedPath, err := expandTilde(dirPath)
	if err != nil {
		return "", fmt.Errorf("failed to expand tilde: %w", err)
	}

	absPath, err := filepath.Abs(expandedPath)
	if err != nil {
		return "", fmt.Errorf("failed to resolve absolute path: %w", err)
	}

	return absPath, nil
}

// validateOrCreateDir validates if a path is a directory or creates it if it doesn't exist
// Returns the absolute path and an error if something goes wrong
func validateOrCreateDir(absPath string) (string, error) {
	// Check if directory already exists
	if info, err := os.Stat(absPath); err == nil {
		if !info.IsDir() {
			return "", fmt.Errorf("'%s' exists but is not a directory", absPath)
		}
		// Directory exists, return the path
		return absPath, nil
	}

	// Create directory with all parent directories (mkdir -p behavior)
	if err := os.MkdirAll(absPath, 0755); err != nil {
		return "", fmt.Errorf("error creating directory: %w", err)
	}

	return absPath, nil
}

// runMkcd performs the mkcd operation and returns the absolute path
func runMkcd(dirPath string) (string, error) {
	absPath, err := resolvePath(dirPath)
	if err != nil {
		return "", err
	}

	resultPath, err := validateOrCreateDir(absPath)
	if err != nil {
		return "", err
	}

	return resultPath, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Creates a directory and changes into it.")
		os.Exit(1)
	}

	dirPath := os.Args[1]
	resultPath, err := runMkcd(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Output the absolute path so the shell can cd into it
	fmt.Println(resultPath)
}
