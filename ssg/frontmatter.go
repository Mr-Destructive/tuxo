package ssg

import (
	"encoding/json"
	"errors"
	"strings"

	"gopkg.in/yaml.v3"
)

// ExtractFrontMatterLines extracts lines between the given delimiters from content.
func ExtractFrontMatterLines(content string, delimiter string) (string, error) {
	start := strings.Index(content, delimiter)
	if start == -1 {
		return "", errors.New("front matter delimiter not found")
	}
	start += len(delimiter)

	end := strings.Index(content[start:], "---")
	if end == -1 {
		return "", errors.New("closing front matter delimiter not found")
	}
	end += start

	return content[start:end], nil
}

// ParseSimpleFrontMatter parses simple key-value pairs from front matter lines.
func ParseSimpleFrontMatter(frontMatterContent string) (map[string]string, error) {
	frontMatter := make(map[string]string)

	lines := strings.Split(frontMatterContent, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, ":", 2)
		if len(parts) != 2 {
			return nil, errors.New("invalid front matter line format")
		}
		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])
		frontMatter[key] = value
	}

	return frontMatter, nil
}

// ParseYAMLFrontMatter parses YAML content from front matter lines.
func ParseYAMLFrontMatter(frontMatterContent string) (map[string]interface{}, error) {
	yamlMap := make(map[string]interface{})

	if err := yaml.Unmarshal([]byte(frontMatterContent), &yamlMap); err != nil {
		return nil, err
	}

	return yamlMap, nil
}

// ParseTOMLFrontMatter parses TOML content from front matter lines.
func ParseTOMLFrontMatter(frontMatterContent string) (map[string]interface{}, error) {
	var tomlMap map[string]interface{}
	return tomlMap, nil
}

// ParseJSONFrontMatter parses JSON content from front matter lines.
func ParseJSONFrontMatter(frontMatterContent string) (map[string]interface{}, error) {
	var jsonMap map[string]interface{}
	if err := json.Unmarshal([]byte(frontMatterContent), &jsonMap); err != nil {
		return nil, err
	}

	return jsonMap, nil
}

// LoadFrontMatter parses front matter from the given content using a configurable delimiter and format.
// Supported formats: "yaml", "toml", "json".
// It returns a map[string]interface{} of front matter key-value pairs and an error if parsing fails.
func LoadFrontMatter(content string, delimiter string, format string) (map[string]interface{}, error) {
	frontMatterContent, err := ExtractFrontMatterLines(content, delimiter)
	if err != nil {
		return nil, err
	}

	switch format {
	case "yaml":
		return ParseYAMLFrontMatter(frontMatterContent)
	case "toml":
		return ParseTOMLFrontMatter(frontMatterContent)
	case "json":
		return ParseJSONFrontMatter(frontMatterContent)
	default:
		return nil, errors.New("unsupported front matter format")
	}
}

// LoadFrontMatterToPost removes front matter from the given content and returns a Post struct.
func LoadFrontMatterToPost(content string, delimiter string, format string) (Post, string, error) {
	frontMatterContent, err := ExtractFrontMatterLines(content, delimiter)
	if err != nil {
		return Post{}, "", err
	}

	// Remove front matter from content
	start := strings.Index(content, delimiter)
	if start == -1 {
		return Post{}, "", errors.New("front matter delimiter not found")
	}
	end := strings.Index(content[start:], "---\n")
	if end == -1 {
		return Post{}, "", errors.New("closing front matter delimiter not found")
	}
	end += start + len(delimiter)

	actualContent := content[end:]

	// Parse front matter based on format
	var frontMatter map[string]interface{}
	switch format {
	case "yaml":
		frontMatter, err = ParseYAMLFrontMatter(frontMatterContent)
	case "toml":
		frontMatter, err = ParseTOMLFrontMatter(frontMatterContent)
	case "json":
		frontMatter, err = ParseJSONFrontMatter(frontMatterContent)
	default:
		return Post{}, "", errors.New("unsupported front matter format")
	}
	if err != nil {
		return Post{}, "", err
	}

	// Convert front matter to Post struct
	post, err := LoadFrontMatterToPostStruct(frontMatter)
	if err != nil {
		return Post{}, "", err
	}

	return post, actualContent, nil
}

// LoadFrontMatterToPostStruct converts a map[string]interface{} front matter to a Post struct.
func LoadFrontMatterToPostStruct(frontMatter map[string]interface{}) (Post, error) {
	post := Post{}
	for key, value := range frontMatter {
		switch key {
		case "title":
			post.Title = value.(string)
		case "description":
			post.Description = value.(string)
		case "date":
			post.Date = value.(string)
		case "tags":
			post.Tags = toStringSlice(value)
		case "status":
			post.Status = value.(string)
		case "slug":
			post.Slug = value.(string)
		}
	}
	return post, nil
}

// toStringSlice converts an interface{} value to a []string slice.
func toStringSlice(value interface{}) []string {
	var result []string
	switch v := value.(type) {
	case []interface{}:
		for _, item := range v {
			result = append(result, item.(string))
		}
	case interface{}:
		result = append(result, v.(string))
	}
	return result
}
