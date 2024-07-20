package ssg

import (
	"bytes"
	"fmt"
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

func WriteHTML(outputDir, filename, feed, content string) error {
	fileORFeed := ""
	if filename == "" {
		fileORFeed = feed + ".html"
		os.WriteFile(filepath.Join(outputDir, fileORFeed), []byte(content), 0644)
	}
	if feed == "" {
		fileORFeed = filename + "/index.html"
		os.MkdirAll(filepath.Join(outputDir, filename), 0755)
		os.WriteFile(filepath.Join(outputDir, fileORFeed), []byte(content), 0644)
	}
	if feed != "" && filename != "" {
		os.MkdirAll(filepath.Join(outputDir, filename), 0755)
		os.WriteFile(filepath.Join(outputDir, filename, feed), []byte(content), 0644)
	}
	fmt.Println(fileORFeed)
	return nil
}
