package main

import (
	"fmt"
	"math"
	"strconv"
)

func ModN(N uint, i int) int {
	m := i % int(N)
	if m < 0 {
		m += int(N)
	}
	return m
}

// OrderMod returns the order of a modulo r
// timing is ____
func OrderMod(a int, r int) int {

	return 0
}

// FastPower (without modulo by N)
func FastPower(g int, A int) int {
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

// EulerTotient function finds the number of numbers less than x that are relatively prime to x
// timing is _____
func EulerTotient(x int) int {

	return 0
}

// PerfectPower determines whether n can be represented as a perfect power a^b
// timing is O(log^3(n))
func PerfectPower(n int) bool {
	var bMax = int(math.Log2(float64(n))) + 1
	for b := 2; b <= bMax; b++ {
		var aMin = 2
		var aMax = int(math.Pow(float64(2), float64(64)/float64(b))) - 1
		if FastPower(b, aMin) == n {
			return true
		}
		if FastPower(b, aMax) == n {
			return true
		}
		for aMax-aMin != 1 {
			//fmt.Println(aMin)
			//fmt.Println(aMin)
			var aMid = (aMin + aMax) / 2
			var nMid = FastPower(aMid, b)
			if nMid == n {
				return true
			} else if nMid < n {
				aMin = aMid
			} else { // nMid > n
				aMax = aMid
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
	var r = 0
	// step 3 -- check GCD between 1 and r

	// step 4 -- if n <= r
	if n <= r {
		return true
	}
	// step 5 -- for loop

	// step 6 -- otherwise return true
	return true
}

func main() {
	//var n = 49
	//var test = PerfectPower(n)
	//if test {
	//	fmt.Println(strconv.Itoa(n) + " is a perfect power")
	//} else {
	//	fmt.Println(strconv.Itoa(n) + " is not a perfect power")
	//}
	for n := 2; n <= 10000; n++ {
		var test = PerfectPower(n)
		if test {
			fmt.Println(strconv.Itoa(n) + " is a perfect power")
		}
	}
}
