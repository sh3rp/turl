package db

import (
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/sh3rp/turl/logger"
	"github.com/sh3rp/turl/proto"
)

var URLS = []byte("urls")

type TurlDB interface {
	Create(proto.URL) (proto.URL, error)
	Get(proto.URL) (proto.URL, error)
	Close() error
}

type boltTurlDB struct {
	db     *bolt.DB
	logger logger.Log
}

func NewDB(dir string, logger logger.Log) (TurlDB, error) {
	bolt, err := bolt.Open(fmt.Sprintf("%s/turl.db", dir), 0600, nil)
	if err != nil {
		return nil, err
	}

	turlDB := boltTurlDB{
		db:     bolt,
		logger: logger,
	}

	return turlDB, nil
}

func (b boltTurlDB) Create(url proto.URL) (proto.URL, error) {
	err := b.db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists(URLS)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(fmt.Sprintf("u-%s", url.Url)), []byte(url.Hash))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(fmt.Sprintf("h-%s", url.Hash)), []byte(url.Url))
		if err != nil {
			return err
		}
		return nil
	})
	return url, err
}

func (b boltTurlDB) Get(url proto.URL) (proto.URL, error) {
	var newUrl *proto.URL
	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket(URLS)

		if url.Hash != "" {
			u := bucket.Get([]byte(fmt.Sprintf("h-%s", url.Hash)))
			newUrl = &proto.URL{Url: string(u), Hash: url.Hash}
		}

		return nil
	})
	return *newUrl, err
}

func (b boltTurlDB) Close() error {
	return b.db.Close()
}
