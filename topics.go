package main

import (
	"fmt"

	"github.com/boltdb/bolt"
)

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
			fmt.Println(markDone, string(name))
			return nil
		})
	})
}

func deleteTopic(topic string) {
	db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte(topic))
		if err != nil {
		}
		fmt.Println("Deleted topic ", topic)
		return nil
	})
}
