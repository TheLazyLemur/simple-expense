package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

var providers = []string{"gmail.com", "yahoo.com", "hotmail.com", "protonmail.com"}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min int64, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString returns a random string of length n from the alphabet
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomUsername returns a random username
func RandomUsername() string {
	return RandomString(8)
}

// RandomEmail returns a random email as well as a random email provider
func RandomEmail() string {
	return RandomString(8) + "@" + providers[rand.Intn(len(providers))]
}
