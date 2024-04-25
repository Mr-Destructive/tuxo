package test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/mr-destructive/tuxo/ssg"
)

func TestLoadPosts(t *testing.T) {
	// Create a temporary directory for test posts
	tmpDir, err := os.MkdirTemp("", "test-posts-")
	if err != nil {
		t.Fatalf("Failed to create temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Create temporary test post file
	testPostFile := filepath.Join(tmpDir, "test-post.md")
	testContent := []byte(`This is the content of the test post.`)
	err = os.WriteFile(testPostFile, testContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test post file: %v", err)
	}

	// Mock TuxoConfig
	config := &ssg.TuxoConfig{PostDir: tmpDir}

	// Test LoadPosts function
	posts, err := ssg.LoadPosts(config)
	if err != nil {
		t.Errorf("LoadPosts returned an error: %v", err)
	}

	// Check if the loaded posts contain the test post
	if len(posts.Posts) != 1 {
		t.Errorf("Expected 1 post, got %d", len(posts.Posts))
	}

	// Check the content of the loaded post
	expectedContent := "This is the content of the test post."
	if posts.Posts[0].Content != expectedContent {
		t.Errorf("Unexpected post content. Expected: %s, Got: %s", expectedContent, posts.Posts[0].Content)
	}
}
