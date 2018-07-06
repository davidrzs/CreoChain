package chain

// // InitializeDatabase ..
// func InitializeDatabase(databaseName string) *bolt.DB {
// 	// Open the my.db data file in your current directory.
// 	// It will be created if it doesn't exist.
// 	db, err := bolt.Open(databaseName, 0600, nil)
// 	fmt.Print(db)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return db
// }
//
// //CreateBucket creates a bucket with the supplied name in the supplied database.
// func CreateBucket(db *bolt.DB, bucket string) error {
// 	err := db.Update(func(tx *bolt.Tx) error {
// 		_, err := tx.CreateBucketIfNotExists([]byte("MyBucket"))
// 		if err != nil {
// 			return fmt.Errorf("create bucket: %s", err)
// 		}
// 		return nil
// 	})
// 	return err
// }
//
// //AddToStore stores a key and value in a bucket in a database. It returns an error if someting goes wrong.
// func AddToStore(db *bolt.DB, bucket string, key string, value string) error {
// 	err := db.Update(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(bucket))
// 		err := b.Put([]byte(key), []byte(value))
// 		return err
// 	})
// 	return err
// }
//
// //RetrieveFromStore retrieves a value for a given key in a bucket in a database. It returns an error if someting goes wrong.
// func RetrieveFromStore(db *bolt.DB, bucket string, key string) (string, error) {
// 	var v string
// 	err := db.View(func(tx *bolt.Tx) error {
// 		b := tx.Bucket([]byte(bucket))
// 		v = string(b.Get([]byte(key)))
// 		return nil
// 	})
// 	return v, err
// }
