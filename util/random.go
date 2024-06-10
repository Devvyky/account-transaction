package util

import (
	"math"
	"math/rand"
	"strings"
)

const numerics = "1234567890"

// RandomString generates a random string of length n
func RandomNumericString(n int) string {
	var sb strings.Builder
	k := len(numerics)

	for i := 0; i < n; i++ {
		c := numerics[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomInt generates a random integer between min and max
func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min+1)
}

// RandomMoney generates a random amount of money
func RandomAmount() float64 {
	amount := randomFloat64(0.00, 1000.00)
	return math.Round(amount*100) / 100
}
