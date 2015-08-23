package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const strike = "L\u0336"

var (
	homeDir  = os.Getenv("HOME")
	dataPath = filepath.Join(homeDir, ".tinker")
)

func createDataPath() {
	// Create data directory.
	err := os.Mkdir(dataPath, 0755)
	if err != nil {
		e := err.(*os.PathError)
		if strings.Contains(e.Err.Error(), "file exists") {
			// Ignore error if dir already exists.
		} else {
			log.Fatal("createDataPath-", err)
		}
	}
}

func strikeText(text string) string {
	var strikeLine string
	for _, c := range text {
		strikeLine += fmt.Sprintf("%c\u0336", c)
	}
	return strikeLine
}
