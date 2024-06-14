package ssg

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func RenderContent(tmpl *template.Template, post Post) (string, error) {
	var templateName string
	fmt.Println(post.Type)
	if post.Type == "posts" {
		templateName = "post.html"
	}
	var rendered bytes.Buffer
	err := tmpl.ExecuteTemplate(&rendered, templateName, post)
	fmt.Println(rendered.String())
	if err != nil {
		return "", err
	}
	return rendered.String(), nil
}

func WriteHTML(outputDir, filename, content string) error {
	outputPath := filepath.Join(outputDir, filename)
	fmt.Println(outputPath)
	err := os.MkdirAll(filepath.Dir(outputPath), os.ModePerm)
	if err != nil {
		return err
	}
	return os.WriteFile(outputPath, []byte(content), 0644)
}
