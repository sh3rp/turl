package hash

type HashProvider interface {
	Hash(string) string
}
