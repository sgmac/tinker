package main

import (
	"fmt"
	"strconv"

	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
)

func startIdea(topic, id string) {
	if !isValidTopic(topic) {
		logrus.Error(ErrorNotExist)
	}

	// Create special bucket to keep track of started ideas.
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("started"))
		if err != nil {
			return err
		}
		return nil
	})

	// Retrive the idea for the given topic and store it
	// in the started bucket.
	idea := getIdea(topic, id)

	t := fmt.Sprintf(" (%s)", topic)
	db.Update(func(tx *bolt.Tx) error {
		nideas := 1
		b := tx.Bucket([]byte("started"))
		if b != nil {
			b.ForEach(func(_ []byte, _ []byte) error {
				nideas++
				return nil
			})
			// Store idead and topic
			return b.Put([]byte(strconv.Itoa(nideas)), []byte(idea+t))
		}
		return ErrorNotExist
	})
}