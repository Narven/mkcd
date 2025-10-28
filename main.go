package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <directory>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "Creates a directory and changes into it.")
		os.Exit(1)
	}

	dirPath := os.Args[1]

	// Expand tilde and handle relative/absolute paths
	if len(dirPath) > 0 && dirPath[0] == '~' {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		if len(dirPath) > 1 {
			dirPath = filepath.Join(home, dirPath[1:])
		} else {
			dirPath = home
		}
	}

	// Get absolute path
	absPath, err := filepath.Abs(dirPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Check if directory already exists
	if info, err := os.Stat(absPath); err == nil {
		if !info.IsDir() {
			fmt.Fprintf(os.Stderr, "Error: '%s' exists but is not a directory\n", absPath)
			os.Exit(1)
		}
		// Directory exists, just cd into it
		fmt.Println(absPath)
		os.Exit(0)
	}

	// Create directory with all parent directories (mkdir -p behavior)
	if err := os.MkdirAll(absPath, 0755); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating directory: %v\n", err)
		os.Exit(1)
	}

	// Output the absolute path so the shell can cd into it
	fmt.Println(absPath)
}
