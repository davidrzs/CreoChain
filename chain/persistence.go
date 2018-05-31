package chain

import (
	"fmt"
	bolt "github.com/coreos/bbolt"
	"log"
)

// InitializeDatabase ..
func InitializeDatabase() *bolt.DB {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	db, err := bolt.Open("my.db", 0600, nil)
	fmt.Print(db)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CreateBucket(db *bolt.DB, bucket string, key string, value string) error {
	return nil
}

//AddToStore stores a key and value in a bucket in a database. It returns an error if someting goes wrong.
func AddToStore(db *bolt.DB, bucket string, key string, value string) error {
	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

//RetrieveFromStore retrieves a value for a given key in a bucket in a database. It returns an error if someting goes wrong.
func RetrieveFromStore(db *bolt.DB, bucket string, key string) (string, error) {
	var v string
	err := db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		v = string(b.Get([]byte(key)))
		return nil
	})
	return v, err
}
