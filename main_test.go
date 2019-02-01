package main

import (
	"os"
	"testing"
)

func TestStitchInit(t *testing.T) {

	conf := "stitchconf.yml"

	StitchInit()

	if _, err := os.Stat("./" + conf); os.IsNotExist(err) {
		t.Error("StitchInit() was unable to create the conf file.")
	}

}
