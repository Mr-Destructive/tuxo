package ssg

import (
	"strings"
)

type Feed struct {
	Type        string
	Title       string
	Description string
	Link        string
	Template    string
	Author      struct {
		Name      string
		Email     string
		Twitter   string
		Facebook  string
		Instagram string
		Youtube   string
		Substack  string
	}
}

func GenerateFeed(config *TuxoConfig) ([]string, error) {
	feeds := []string{}
	config.Feeds["default"] = Feed{
		Type:  "post",
		Title: "index",
	}

	for _, feed := range config.Feeds {
		items, err := getFeedItems(feed, config.PostDir)
		if err != nil {
			return nil, err
		}

		if feed.Template == "" {
			feed.Template = "feed.html"
		}

		tmpl, err := LoadTemplates(config)
		if err != nil {
			return nil, err
		}
		for i, item := range items {
			item, _, err = LoadFrontMatterToPost(item.Content, "---json", config.Type)
			if err != nil {
				return nil, err
			}
			if item.Slug == "" {
				item.Slug = item.Title
			}
			items[i] = item
		}

		payload := map[string]interface{}{
			"Title": feed.Title,
			"Posts": items,
		}

		rendered := new(strings.Builder)

		err = tmpl.ExecuteTemplate(rendered, feed.Template, payload)
		if err != nil {
			return nil, err
		}
		feedName := ""
		fileName := ""
		if feed.Title == "index" {
			fileName = ""
		} else {
			fileName = feed.Title
			feedName = "index.html"
		}
		err = WriteHTML(config.OutputDir, fileName, feedName, rendered.String())
		if err != nil {
			return nil, err
		}
		feeds = append(feeds, rendered.String())
	}

	return feeds, nil
}

/*
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
		log.Println(len(items))

		sb.WriteString("<ul>\n")
		for _, item := range items {
			post, _, err := LoadFrontMatterToPost(item.Content, "---json", config.Type)
			if err != nil {
				return nil, err
			}
			if post.Slug == "" {
				post.Slug = post.Title
			}
			fmt.Println(post)
			sb.WriteString(fmt.Sprintf("<li><a href=\"%s\">%s</a></li>\n", post.Slug, post.Title))
		}
		sb.WriteString("</ul>\n")

		feedContents = append(feedContents, sb.String())
		WriteHTML(config.OutputDir, "", feed.Title, sb.String())
	}
*/
//
//	return feedContents, nil
//}

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
