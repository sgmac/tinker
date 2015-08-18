package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
