package php

import (
	"crypto/rand"
)

// RandomBytes generates cryptographically secure pseudo-random bytes
func RandomBytes(length int) []byte {
	b := make([]byte, length)
	rand.Read(b)
	return b
}

// RandomString generates cryptographically secure pseudo-random bytes and return its hexadecimal representation
func RandomString(length int) string {
	return Bin2hex(RandomBytes(length))
}
