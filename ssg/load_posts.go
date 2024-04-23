package ssg

type Post struct {
	Type        string
	Title       string
	Description string
	Date        string
	Tags        []string
	Status      string
	Slug        string
}

type Posts struct {
	Posts    []Post
	Type     string
	BasePath string
}

func LoadPosts() (Posts, error) {
    return Posts{}, nil
}
