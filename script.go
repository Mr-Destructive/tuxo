package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

type Post struct {
	ID          int
	Title       string
	Description string
	Content     string
	Author      string
	CreatedAt   string
	UpdatedAt   string
}

func fetchPosts(db *sql.DB) ([]Post, error) {
	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []Post
	for rows.Next() {
		var post Post
		postVal := reflect.ValueOf(&post).Elem()
		var scanArgs []interface{}
		for i := 0; i < postVal.NumField(); i++ {
			scanArgs = append(scanArgs, postVal.Field(i).Addr().Interface())
		}
		if err := rows.Scan(scanArgs...); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func generateMarkdown(posts []Post) error {
	for _, post := range posts {
		// create posts directory if it doesn't exist
		if _, err := os.Stat("posts"); os.IsNotExist(err) {
			if err := os.Mkdir("posts", 0755); err != nil {
				return err
			}
		}
		filePath := fmt.Sprintf("posts/%d.md", post.ID)
		file, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer file.Close()
		frontmatter := "---\n"
		tempPost := reflect.ValueOf(post)
		if tempPost.Kind() == reflect.Ptr {
			tempPost = reflect.Indirect(tempPost)
		}
		if tempPost.Kind() != reflect.Struct {
			return fmt.Errorf("expected struct, got %s", tempPost.Kind())
		}
		n := tempPost.NumField()
		for i := 0; i < n; i++ {
			field := tempPost.Field(i)
			fieldName := tempPost.Type().Field(i).Name
			fieldValue := field.String()

			if field.Kind() == reflect.String && fieldValue != "" && fieldName != "Content" {
				frontmatter += fmt.Sprintf("%s: %s\n", fieldName, fieldValue)
			}
		}
		content := fmt.Sprintf("%s---\n\n%s", frontmatter, post.Content)
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
		log.Fatal(err)
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
