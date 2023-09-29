package utils

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"
const alphaNumeric = alphabet + "0123456789"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func randomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

func randomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func randomPassword(n int) string {
	var sb strings.Builder
	k := len(alphaNumeric)
	for i := 0; i < n; i++ {
		c := alphaNumeric[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()

}

func RandomOwner() string {
	return randomString(6)
}

func RandomPassword() string {
	return randomPassword(20)
}

func RandomString(n int) string {
	return randomString(n)
}

func RandomCurrency() string {
	currencies := []string{"EUR", "USD", "CAD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}
