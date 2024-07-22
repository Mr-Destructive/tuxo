package ssg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// LoadStatic recursively copies files and directories from staticDir to outputDir
func LoadStatic(config *TuxoConfig) error {
	staticDir := config.StaticDir
	outputDir := config.OutputDir

	if staticDir == "" {
		return fmt.Errorf("no static dir provided")
	}

	if _, err := os.Stat(staticDir); err != nil {
		return err
	}

	err := filepath.Walk(staticDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get relative path to the file or directory
		relativePath, err := filepath.Rel(staticDir, path)
		if err != nil {
			return err
		}

		// Destination path
		destPath := filepath.Join(outputDir, relativePath)

		if info.IsDir() {
			// Create directory if it doesn't exist
			err := os.MkdirAll(destPath, 0755)
			if err != nil {
				return err
			}
		} else {
			// Copy file contents
			err := copyFile(path, destPath)
			if err != nil {
				return err
			}
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}
