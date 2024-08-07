package ssg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type TuxoConfig struct {
	Type        string          `yaml:"type" json:"type" toml:"type"`
	PostDir     string          `yaml:"post_dir" json:"post_dir" toml:"post_dir"`
	StaticDir      string          `yaml:"static_dir" json:"static_dir" toml:"static_dir"`
	TemplateDir string          `yaml:"template_dir" json:"template_dir" toml:"template_dir"`
	OutputDir   string          `yaml:"output_dir" json:"output_dir" toml:"output_dir"`
	Feeds       map[string]Feed `yaml:"feeds" json:"feeds" toml:"feeds"`
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
		jsonConfig := TuxoConfig{
			Type: "json",
		}
		err = json.NewDecoder(file).Decode(&jsonConfig)
		if err != nil {
			return nil, err
		}
		return &jsonConfig, nil
	case ".yaml", ".yml":
		yamlConfig := TuxoConfig{
			Type: "yaml",
		}
		err := yaml.NewDecoder(file).Decode(&yamlConfig)
		if err != nil {
			return nil, err
		}
		return &yamlConfig, nil
	case ".toml":
		return nil, nil
	}
	return nil, nil
}
