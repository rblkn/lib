package adder

import "golang.org/x/exp/constraints"

func Add[T constraints.Integer](a, b T) T {
	return a + b
}
