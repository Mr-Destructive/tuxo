package main

import (
	"flag"
	"fmt"

	"github.com/mr-destructive/tuxo/ssg"
)

func main() {
	fmt.Println("TUXO Init")
	configFiles, err := ssg.LoadConfigFilePath()
	if err != nil {
		panic(err)
	}
	configFile := configFiles[0]
	config, err := ssg.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}
	post, err := ssg.LoadPosts(config)
	if err != nil {
		panic(err)
	}
	templates, err := ssg.LoadTemplates(config)
	if err != nil {
		panic(err)
	}
	for _, post := range post.Posts {
		post, parsedHTML, err := ssg.RenderContent(templates, post, config)
		if err != nil {
			panic(err)
		}
		if post.Slug == "" {
			post.Slug = post.Title
		}
		err = ssg.WriteHTML(config.OutputDir, post.Slug, "", parsedHTML)
		if err != nil {
			panic(err)
		}
	}
	//load feeds
	_, err = ssg.GenerateFeedHTML(config)
	if err != nil {
		panic(err)
	}
	// copy static files
	err = ssg.LoadStatic(config)
	var portFlag = flag.Int("p", 8080, "port number")
	flag.Parse()
	ssg.StartServer(*portFlag, config.OutputDir)
}
