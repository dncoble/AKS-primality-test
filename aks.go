package aks

import (
	"math"
)

func ModN(N uint, i int) int {
	m := i % int(N)
	if m < 0 {
		m += int(N)
	}
	return m
}

// FastPower (without modulo by N)
func FastPower(g int, A uint) int {
	var b int
	a := g
	b = 1
	for A > 0 {
		if A%2 == 1 {
			b = b * a
		}
		a = a * a
		A = A / 2
	}
	return b
}

// PerfectPower determines whether n can be represented as a perfect power a^b
// by searching for timing is O(log^3(n))
func PerfectPower(n int) bool {
	var bMax = int(math.Log2(float64(n))) + 1
	for b := 2; b <= bMax; b++ {
		var aMin uint = 2
		var aMax uint = uint(n)
		var aMid uint = (aMin + aMax) / 2
		for aMax-aMin != 1 {
			var nMid uint = FastPower(n, b, aMid)
			if (nMid == n) {
				return
			}
		}
	}

	return false
}

// AKS algorithm false = composite, true = prime
func AKS(n int) bool {
	// step 1 -- find if n is a perfect power
	if PerfectPower(n) {
		return false
	}
	// step 2 -- find r
	var r int = 0
	// step 3 -- check GCD between 1 and r

	// step 4 -- if n <= r
	if n <= r {
		return true
	}
	// step 5 -- for loop

	// step 6 -- otherwise return true
	return true
}
