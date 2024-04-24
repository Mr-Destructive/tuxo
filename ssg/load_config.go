package ssg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type TuxoConfig struct {
	PostDir string `yaml:"post_dir" json:"post_dir" toml:"post_dir"`
	Static  string `yaml:"static" json:"static" toml:"static"`
}

func LoadConfigFilePath() ([]string, error) {
	fileNamePattern := "tuxo"
	allowedExtensions := []string{".yaml", ".yml", ".json", ".toml"}
	configFiles := []string{}
	fileList := []os.FileInfo{}
	for _, ext := range allowedExtensions {
		if configFile, err := filepath.Glob(fileNamePattern + ext); err == nil {
			if len(configFile) > 0 {
				configFiles = append(configFiles, configFile[0])
				fileObj, err := os.Stat(configFile[0])
				if err != nil {
					return nil, err
				}
				fileList = append(fileList, fileObj)
			}
		} else {
			return nil, err
		}
	}
	if len(configFiles) == 0 {
		return nil, fmt.Errorf("no config file found")
	}
	if len(configFiles) > 1 {
		lastModTime := fileList[0].ModTime()
		configFiles = []string{fileList[0].Name()}
		for _, file := range fileList {
			if file.ModTime().After(lastModTime) {
				lastModTime = file.ModTime()
				configFiles = []string{file.Name()}
			}
		}
		return configFiles, nil
		//return nil, fmt.Errorf("multiple config files found")
	}
	return configFiles, nil
}

func LoadConfig(configPath string) (*TuxoConfig, error) {
	if configPath == "" {
		return nil, fmt.Errorf("no config file path provided")
	}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	switch filepath.Ext(configPath) {
	case ".json":
		jsonConfig := TuxoConfig{}
		err = json.NewDecoder(file).Decode(&jsonConfig)
		if err != nil {
			return nil, err
		}
		return &jsonConfig, nil
	case ".yaml", ".yml":
		return nil, nil
	case ".toml":
		return nil, nil
	}
	return nil, nil
}
