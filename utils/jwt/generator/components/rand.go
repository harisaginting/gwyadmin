package components

import (
	"math/rand"
	"time"
)

const charset = "0123456789" + "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

// StringWithCharset creates a random string on 5 characters. a-z A-Z 0-9
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandString returns the String for the Sequence Number used in the identifier.
func RandString() string {
	return StringWithCharset(6, charset)
}
