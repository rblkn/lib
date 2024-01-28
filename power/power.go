package power

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Pow[T constraints.Float](base, exponent T) float64 {
	return math.Pow(float64(base), float64(exponent))
}
