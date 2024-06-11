package util

import (
	"math"
	"math/rand"
	"strings"
)

const numerics = "1234567890"
const alphabet = "abcdefghijklmnopqrstuvwxyz"

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNumericString generates a random string of length n
func RandomNumericString(n int) string {
	var sb strings.Builder
	k := len(numerics)

	for i := 0; i < n; i++ {
		c := numerics[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomFloat64 generates a random float value between min and max
func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min+1)
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomMoney generates a random amount of money
func RandomAmount() float64 {
	amount := RandomFloat64(0.00, 1000.00)
	return math.Round(amount*100) / 100
}
