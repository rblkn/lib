package calc

import (
	"github.com/rblkn/lib/adder"
	"github.com/rblkn/lib/power"
)

func PrintCalc() {
	println(adder.Add(3, 4))
	println(power.Pow(2, 3))
}
