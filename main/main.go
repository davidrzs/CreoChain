package main

import (
	"../chain"
	//"../server"
	"fmt"
	bolt "github.com/coreos/bbolt"
)

func main() {
	db := chain.InitializeDatabase()
	defer db.Close() //remember to close it at the end
	//server.Serve()
	fmt.Println("Up and running")
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		err := b.Put([]byte("answer"), []byte("42"))
		fmt.Print("in here")
		return err
	})

	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("MyBucket"))
		v := b.Get([]byte("answer"))
		fmt.Printf("The answer is: %s\n", v)
		return nil
	})

	fmt.Println("hello")

}
