package main

import (
	"fmt"
	"github.com/mr-destructive/tuxo/ssg"
)

func main() {
	fmt.Println("TUXO Init")
	configFiles, err := ssg.LoadConfigFilePath()
	if err != nil {
		panic(err)
	}
	fmt.Println(configFiles)
	configFile := configFiles[0]
	config, err := ssg.LoadConfig(configFile)
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
	post, err := ssg.LoadPosts(config)
	if err != nil {
		panic(err)
	}
	templates, err := ssg.LoadTemplates(config)
	if err != nil {
		panic(err)
	}
	for _, post := range post.Posts {
		parsedHTML, err := ssg.RenderContent(templates, post)
		if err != nil {
			panic(err)
		}
        fmt.Println(post.Slug)
		err = ssg.WriteHTML(config.OutputDir, post.Slug+".html", parsedHTML)
		if err != nil {
			panic(err)
		}
	}
}
