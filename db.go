package main

import (
	"fmt"
	"log"
	"path/filepath"

	"github.com/boltdb/bolt"
)

var db *bolt.DB

func init() {
	var err error
	dbPath := filepath.Join(dataPath, "bolt.db")
	db, err = bolt.Open(dbPath, 0644, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// A new topic is mapped to a bucket.
func addNewTopic(topic string) {
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(topic))
		if err != nil {
			return err
		}
		return nil
	})
}

func listTopics() {
	db.View(func(tx *bolt.Tx) error {
		return tx.ForEach(func(name []byte, _ *bolt.Bucket) error {
			fmt.Println(string(name))
			return nil
		})
	})
}
