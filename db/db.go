package db

import (
	"github.com/boltdb/bolt"
	"github.com/sh3rp/turl/proto"
)

type TurlDB interface {
	Create(proto.URL) (proto.URL, error)
	Get(proto.URL) (proto.URL, error)
}

type boltTurlDB struct {
	db *bolt.DB
}
