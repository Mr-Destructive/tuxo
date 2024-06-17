package ssg

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func RenderContent(tmpl *template.Template, post Post, config *TuxoConfig) (Post, string, error) {
	var postObj Post
	var templateName string
	if post.Type == "post" {
		templateName = "post.html"
	}
	var rendered bytes.Buffer
	post, content, err := LoadFrontMatterToPost(post.Content, "---json", config.Type)
	post.Content = content
	err = tmpl.ExecuteTemplate(&rendered, templateName, post)
	if err != nil {
		return postObj, "", err
	}
	return post, rendered.String(), nil
}

func WriteHTML(outputDir, filename, content string) error {
	err := os.MkdirAll(filepath.Join(outputDir, filename), 0755)
	err = os.WriteFile(filepath.Join(outputDir, filename, "index.html"), []byte(content), 0644)
	if err != nil {
		return err
	}
	return nil
}
