package hash

import (
	"encoding/hex"
	"hash/adler32"
)

type adler32Hash struct{}

func NewAdler32Provider() HashProvider {
	return adler32Hash{}
}

func (a adler32Hash) Hash(str string) string {
	hasher := adler32.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}
