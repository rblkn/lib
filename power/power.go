package power

import (
	"math"

	"golang.org/x/exp/constraints"
)

// Pow returns base^exponent. It is a generic function that can be used with
func Pow[T constraints.Float](base, exponent T) float64 {
	return math.Pow(float64(base), float64(exponent))
}
