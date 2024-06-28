package ssg

import (
	"fmt"
	"strings"
)

type Feed struct {
	Type        string
	Title       string
	Description string
	Link        string
	Author      struct {
		Name  string
		Email string
		Link  string
	}
}

func GenerateFeed(config *TuxoConfig) ([]string, error) {
	var feedContents []string

	for _, feed := range config.Feeds {
		var sb strings.Builder
		// parse feed interface into a map of string and interface

		sb.WriteString(fmt.Sprintf("<h1>%s</h1>\n", feed.Title))
		sb.WriteString(fmt.Sprintf("<p>%s</p>\n", feed.Description))
		sb.WriteString(fmt.Sprintf("<a href=\"%s\">Website</a>\n", feed.Link))
		sb.WriteString(fmt.Sprintf("<p>Author: <a href=\"%s\">%s</a> (%s)</p>\n", feed.Author.Link, feed.Author.Name, feed.Author.Email))

		// Assuming you have a function to get the items for the feed
		items, err := getFeedItems(feed, config.PostDir)
		if err != nil {
			return nil, err
		}

		sb.WriteString("<ul>\n")
		for _, item := range items {
			sb.WriteString(fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n", item.Slug, item.Title))
		}
		sb.WriteString("</ul>\n")

		feedContents = append(feedContents, sb.String())
	}

	return feedContents, nil
}

func getFeedPostByType(config *TuxoConfig, postType string, allPosts []Post) ([]Post, error) {
	posts := []Post{}
	for _, post := range allPosts {
		if post.Type == postType {
			posts = append(posts, post)
		}
	}
	return posts, nil
}

func getFeedItems(feed Feed, postDir string) ([]Post, error) {
	allPosts, err := LoadPosts(&TuxoConfig{
		PostDir: postDir,
	})
	if err != nil {
		return nil, err
	}
	posts, err := getFeedPostByType(&TuxoConfig{
		PostDir: postDir,
	}, feed.Type, allPosts.Posts)
	if err != nil {
		return nil, err
	}
	return posts, nil
}

func GenerateFeedHTML(config *TuxoConfig) (string, error) {
	feedContents, err := GenerateFeed(config)
	if err != nil {
		return "", err
	}
	return strings.Join(feedContents, "\n"), nil
}
