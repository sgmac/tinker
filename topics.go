package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/boltdb/bolt"
	"gopkg.in/yaml.v2"
)

var defaultTopicFile string

type Topic struct {
	Default string
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

func setDefaultTopic(topic string) {
	defaultTopicFile := filepath.Join(dataPath, "topic")
	if _, err := os.Stat(defaultTopicFile); os.IsNotExist(err) {
		data := new(Topic)
		emptyConf, err := yaml.Marshal(data)
		if err != nil {
			logrus.Fatal(err)
		}
		ioutil.WriteFile(defaultTopicFile, emptyConf, 0644)
	}
	st := new(Topic)
	data, err := ioutil.ReadFile(defaultTopicFile)
	if err != nil {
		logrus.Fatal(err)
	}
	// Read the previous default topic
	err = yaml.Unmarshal(data, &st)
	if err != nil {
		logrus.Fatal(err)
	}

	// Write new topic
	st.Default = topic
	newData, err := yaml.Marshal(st)
	ioutil.WriteFile(defaultTopicFile, newData, 0644)

}
