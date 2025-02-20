package db

import (
	"errors"
	"fmt"
	"os"
	"path"
	"time"

	"github.com/CosmicPredator/chibi/internal"
	bolt "go.etcd.io/bbolt"
)

// define DbContext type to make CRUD operations
type DbContext struct {}

// creates a new bucket and if exists, it'll skip
func (dc DbContext) createBucket() error {
    db, err := dc.openDbConn()
    defer db.Close()
    if err != nil {
        return err
    }
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(internal.BOLT_BUCKET_NAME))
		if err != nil {
			return err
		}
		return nil
	})
}

// cleans up and prepares os sepcific config folder
// to save .db file
func (dc DbContext) InitDB() error {
	osConfigPath, _ := os.UserConfigDir()
	configDir := path.Join(osConfigPath, "chibi")
	_, err := os.Stat(configDir)

	if err == nil {
		os.RemoveAll(configDir)
	}
	os.MkdirAll(configDir, 0755)
	return dc.createBucket()
}

// opens the .db file and returns the instance
func (dc *DbContext) openDbConn() (*bolt.DB, error) {
	osConfigDir, _ := os.UserConfigDir()
	configFilePath := path.Join(osConfigDir, "chibi", internal.BOLT_DB_NAME)

	db, err := bolt.Open(configFilePath, 0755, &bolt.Options{
        Timeout: 5 * time.Second,
    })
	if err != nil {
		return nil, err
	}
	return db, nil
}

// writes the key value pair to db
func (dc DbContext) Set(key string, value string) error {
    db, err := dc.openDbConn()
    defer db.Close()
    if err != nil {
        return err
    }
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BOLT_BUCKET_NAME))
		if b == nil {
			return errors.New("DB Bucket does not exist")
		}
		err := b.Put([]byte(key), []byte(value))
		return err
	})
	return err
}

// reads the value of specified key from db
func (dc DbContext) Get(key string) (string, error) {
    db, err := dc.openDbConn()
    defer db.Close()
    if err != nil {
        return "", err
    }

	var value string
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(internal.BOLT_BUCKET_NAME))
		if b == nil {
			return errors.New("DB Bucket does not exist")
		}
		value = string(b.Get([]byte(key)))
		if value == "" {
			return fmt.Errorf("no value found for the key %s", key)
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return value, err
}

func NewDbConn() *DbContext {
	return &DbContext{}
}
