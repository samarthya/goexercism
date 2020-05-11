package cipher

// Cipher Interfaces are named collections of method signatures.
type Cipher interface {
	Encode(string) string
	Decode(string) string
}
