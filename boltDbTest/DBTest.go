package main

import (
	"fmt"
	"log"
	"time"

	"github.com/boltdb/bolt"
)

type Appointment struct {
	Username string
	Date     time.Time
}

var appoint2 []Appointment

func booking() {
	t, _ := time.Parse("2006-01-02", "2022-10-19")
	appoint2 = append(appoint2, Appointment{Username: "Wilson", Date: t}, Appointment{Username: "Tony", Date: t})
}

func main() {
	// Create a database named "DEMO.db" in your current directory.
	// It will be created if it doesn't exist.
	// And keep it opened.
	db, err := bolt.Open("DEMO.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Close DEMO.db when this main() finished.
	defer db.Close()
	// Create a bucket(table) named "appointments".
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Appointments"))
		if err != nil {
			return fmt.Errorf("create bucket err: %s", err)
		}
		return nil
	})

	// Restore all appointments from "DEMO.db/Appointments"
	db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("Appointments"))

		b.ForEach(func(k, v []byte) error {
			fmt.Printf("key=%s, value=%s\n", k, v)
			return nil
		})
		return nil
	})
	// Book two sample dates.
	booking()

	// Store all appointments to "DEMO.db/Appointments"
	for _, a := range appoint2 {
		db.Update(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte("Appointments"))
			err := b.Put([]byte(a.Username), []byte(a.Date.Format("2006-01-02")))
			return err
		})
	}
}
