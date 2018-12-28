package hash

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var TEST_STRING = "test string"
var TEST_STRING_HASH = "1ac00478"

func Test_Adler32(t *testing.T) {
	provider := adler32Hash{}
	hash := provider.Hash("test string")
	fmt.Printf("%s\n", hash)
	assert.Equal(t, TEST_STRING_HASH, hash)
}
