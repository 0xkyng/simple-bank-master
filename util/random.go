package util

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "abcdefghijklmnopqrstuvwxyz"


// RandomInt generates a random integer btween min & max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min + 1) 
	// returns a random interger
	// Between min and max
}

// RandmomString generates a random string of length n
func RandmomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	// Use a for loop to generate n random characters
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandmomString(6)
}

// RandomMoney generates a random ammount of money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency generates a random currency code
func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "NGN"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}