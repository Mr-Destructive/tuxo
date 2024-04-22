package ssg

import (
	"fmt"
	"path/filepath"
)

func LoadConfig() ([]string, error) {
	fileNamePattern := "turxgo"
	allowedExtensions := []string{".yaml", ".yml", ".json", ".toml"}
	configFiles := []string{}
	for _, ext := range allowedExtensions {
		if configFile, err := filepath.Glob(fileNamePattern + ext); err == nil {
			if len(configFile) > 0 {
				configFiles = append(configFiles, configFile[0])
			}
		} else {
			return nil, err
		}
	}
	if len(configFiles) == 0 {
		return nil, fmt.Errorf("no config file found")
	}
	if len(configFiles) > 1 {
		return nil, fmt.Errorf("multiple config files found")
	}
	return configFiles, nil
}
