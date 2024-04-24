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
}
