package ssg

import (
	"html/template"
	"os"
	"path/filepath"
)

func LoadTemplates(config *TuxoConfig) (*template.Template, error) {
	templateDir := config.TemplateDir
	tmpl := template.New("")
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(path) == ".html" {
			_, err := tmpl.ParseFiles(path)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return tmpl, err
}
