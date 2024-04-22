package test

import (
	"testing"

	"github.com/mr-destructive/turxgo/ssg"
)

func TestLoadConfig(t *testing.T) {
	configFiles, err := ssg.LoadConfig()
	if err != nil {
		t.Error(err)
	}
	if len(configFiles) == 0 {
		t.Error("no config file found")
	}
	if len(configFiles) > 1 {
		t.Error("multiple config files found")
	}
}
