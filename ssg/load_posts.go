package ssg

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Post struct {
	Type        string
	Title       string
	Description string
	Date        string
	Tags        []string
	Status      string
	Slug        string
	Content     string
}

type Posts struct {
	Posts    []Post
	Type     string
	BasePath string
}

func LoadPosts(config *TuxoConfig) (Posts, error) {
	postDir := config.PostDir
	posts := Posts{Type: "posts", BasePath: postDir}

	if postDir == "" {
		return Posts{}, fmt.Errorf("no post dir provided")
	}

	if _, err := os.Stat(postDir); err != nil {
		return Posts{}, err
	}

	files, err := os.ReadDir(postDir)
	if err != nil {
		return Posts{}, err
	}
	for _, file := range files {

		if file.IsDir() {
			continue
		}
		if filepath.Ext(file.Name()) != ".md" {
			continue
		}

		postPath := filepath.Join(postDir, file.Name())
		post, err := readPost(postPath)
		if err != nil {
			return Posts{}, err
		}
		posts.Posts = append(posts.Posts, post)
	}
	return posts, nil
}

func readPost(postPath string) (Post, error) {
	file, err := os.Open(postPath)
	if err != nil {
		return Post{}, err
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		return Post{}, err
	}
	post := Post{
		Title:   postPath,
		Content: string(content),
		Type:    "posts",
        Slug:    filepath.Base(postPath),
	}

	return post, nil
}
