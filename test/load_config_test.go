package test

import (
	"log"
	"testing"

	"github.com/mr-destructive/tuxo/ssg"
)

func TestLoadConfig(t *testing.T) {
	configFiles, err := ssg.LoadConfigFilePath()
	log.Println(configFiles)
	if err != nil {
		t.Error(err)
	}
	if len(configFiles) == 0 {
		t.Error("no config file found")
	}
	if len(configFiles) > 1 {
	}
}
