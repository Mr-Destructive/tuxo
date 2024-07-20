package ssg

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadStatic(config *TuxoConfig) error {
	staticDir := config.StaticDir
	outputDir := config.OutputDir

	if staticDir == "" {
		return fmt.Errorf("no static dir provided")
	}

	if _, err := os.Stat(staticDir); err != nil {
		return err
	}

	files, err := os.ReadDir(staticDir)
	if err != nil {
		return err
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		//create file to outputDir
		os.Create(filepath.Join(outputDir, file.Name()))
		//copy the contents
		contents, err := os.ReadFile(filepath.Join(staticDir, file.Name()))
		if err != nil {
			return err
		}
		err = os.WriteFile(filepath.Join(outputDir, file.Name()), []byte(contents), 0644)
		if err != nil {
			return err
		}

	}
	return nil
}
