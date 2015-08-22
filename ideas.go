package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/boltdb/bolt"
)

var ErrorNotExist = errors.New("topic does not exist")

func addIdea(topic, idea string) {
	err := db.Update(func(tx *bolt.Tx) error {
		nideas := 1
		b := tx.Bucket([]byte(topic))
		if b != nil {
			fmt.Printf("Idea: %s Topic: %s\n", topic, idea)
			b.ForEach(func(_ []byte, _ []byte) error {
				nideas++
				return nil
			})
			return b.Put([]byte(strconv.Itoa(nideas)), []byte(idea))
		}
		return ErrorNotExist
	})
	if err != nil {
		fmt.Println(topic+":", err)
	}
}

func listIdeas(topic string) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(topic))
		if b != nil {
			fmt.Println("<" + topic + ">")
			b.ForEach(func(k []byte, v []byte) error {
				fmt.Printf("%s: %s %s\n", k, check, v)
				return nil
			})
			return nil
		}
		return ErrorNotExist
	})
	if err != nil {
		fmt.Println(topic+":", err)
	}
}

func deleteIdea(topic, id string) {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(topic))
		if b != nil {
			b.Delete([]byte(id))
			return nil
		}
		return ErrorNotExist
	})
	if err != nil {
		fmt.Println(topic+":", err)
	}
}

func getIdea(topic, id string) string {
	var idea string
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(topic))
		if b != nil {
			idea = string(b.Get([]byte(id)))
			return nil
		}
		return ErrorNotExist
	})
	if err != nil {
		fmt.Println(topic+":", err)
	}
	return idea
}
