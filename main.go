package main

import (
	"fmt"
	"github.com/mr-destructive/turxgo/ssg"
)

func main() {
	fmt.Println("TURXGO Init")
	configFiles, err := ssg.LoadConfig()
	if err != nil {
		panic(err)
	}
	fmt.Println(configFiles)
}
