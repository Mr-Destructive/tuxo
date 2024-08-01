package main

import (
	"bufio"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type DBPost struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Content     string `json:"content"`
	Author      string `json:"author"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Post struct {
	DBPost
	Slug string `json:"slug"`
	Type string `json:"type"`
}

func fetchPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post DBPost
		postVal := reflect.ValueOf(&post).Elem()
		var scanArgs []interface{}
		for i := 0; i < postVal.NumField(); i++ {
			scanArgs = append(scanArgs, postVal.Field(i).Addr().Interface())
		}
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		p := Post{DBPost: post, Type: "posts"}
		posts = append(posts, p)
	}
	return posts, nil
}

func generateMarkdown(posts []Post) error {
	for _, post := range posts {
		// create posts directory if it doesn't exist
		if _, err := os.Stat("blog"); os.IsNotExist(err) {
			if err := os.Mkdir("blog", 0755); err != nil {
				return err
			}
		}
		filePath := fmt.Sprintf("blog/%d.md", post.ID)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		jsonKV, err := json.Marshal(post)
		content := fmt.Sprintf("---json{\n%s\n}\n---\n\n%s", string(jsonKV)[1:len(string(jsonKV))-1], post.Content)
		if _, err := file.WriteString(content); err != nil {
			return err
		}
	}
	return nil
}

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		os.Setenv(key, value)
	}

	return scanner.Err()
}
func main() {
	err := LoadEnv(".env")
	if err != nil {
	}
	dbToken := os.Getenv("DB_TOKEN")
	dbName := os.Getenv("DB_NAME")
	dbOrgName := os.Getenv("DB_ORG_NAME")

	url := fmt.Sprintf("libsql://%s-%s.turso.io?authToken=%s", dbName, dbOrgName, dbToken)
	db, err := sql.Open("libsql", url)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	posts, err := fetchPosts(db)
	if err != nil {
		log.Fatalf("Error fetching posts: %v", err)
	}

	if err := generateMarkdown(posts); err != nil {
		log.Fatalf("Error generating markdown: %v", err)
	}
}
