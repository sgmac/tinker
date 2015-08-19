package main

import (
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

const markDone = "\u2713"
const check = "\u09E9"

func init() {
	var err error
	dbPath := filepath.Join(dataPath, "bolt.db")
	db, err = bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
}
