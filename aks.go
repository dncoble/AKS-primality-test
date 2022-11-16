package main

import (
	"math"
	//"strconv"
)

func ModN(N uint, i int) int {
	m := i % int(N)
	if m < 0 {
		m += int(N)
	}
	return m
}

func GCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}

	var na, nb bool
	if a < 0 {
		na = true
		a *= -1
	}
	if b < 0 {
		nb = true
		b *= -1
	}

	var g, u, v int
	u = 1
	g = a
	x := 0
	y := b
	for y != 0 {
		t := ModN(uint(y), g)
		q := g / y
		s := u - q*x
		u = x
		g = y
		x = s
		y = t
	}
	v = (g - a*u) / b

	if !na && !nb {
		return g, u, v
	} else if !na && nb {
		return g, u, -v
	} else if na && !nb {
		return g, -u, v
	} else {
		return g, -u, -v
	}
}

// OrderMod returns the order of a modulo r
// timing is ____
func OrderMod(a int, r int) int {
	gcd, _, _ := GCD(a, r)
	if gcd != 1 {
		return 0
	}
	var i int = 1
	for {
		if FastPowerMod(uint(r), a, uint(i)) == 1 {
			return i
		}
		i++
	}
}

type Polynomial struct {
	d     int   //order
	coefs []int // coefficients
}

// PolynomialMultiply multiplies two polynomials represented by Polynomial struct
// timing is product of orders of the polynomials
func PolynomialMultiply(x, y Polynomial) Polynomial {
	dx := x.d
	dy := y.d
	coefs := make([]int, dx*dy+1)
	for i := 0; i < len(coefs); i++ {
		coefs[i] = 0
	}
	for i, c1 := range x.coefs {
		for j, c2 := range y.coefs {
			coefs[i+j] += c1 * c2
		}
	}
	return Polynomial{dx * dy, coefs}
}

// PolynomialMod x mod y using polynomial division algorithm
// timing is ___
func PolynomialMod(x, y Polynomial) Polynomial {
	return Polynomial{1, []int{1}}
}

// PolynomialFastPower x^n mod(y, N) with x, y polynomials and n, N integers
// timing is ____
func PolynomialFastPower(x Polynomial, n int, y Polynomial, N int) Polynomial {
	var b Polynomial
	a := x
	b = Polynomial{1, []int{1}}
	for n > 0 {
		if n%2 == 1 {
			b = PolynomialMod(PolynomialMultiply(b, a), y)
		}
		a = PolynomialMultiply(a, a)
		n = n / 2
	}
	return b
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

func FastPowerMod(N uint, g int, A uint) int {
	var b int
	a := g
	b = 1
	if A < 0 {
		A = -A
	}
	for A > 0 {
		if A%2 == 1 {
			b = ModN(N, b*a)
		}
		a = ModN(N, a*a)
		A = A / 2
	}
	return b
}

func StepTwo(n int) int {
	var lower int = int(math.Log2(float64(n)) * math.Log2(float64(n)))
	var r int = 2
	for {
		if OrderMod(n, r) > lower {
			return r
		}
		r++
	}
}

// EulerTotient function finds the number of numbers less than x that are relatively prime to x
// timing is _____
func EulerTotient(x int) int {
	var y int = 0
	for i := 0; i < x; i++ {
		gcd, _, _ := GCD(i, x)
		if gcd == 1 {
			y++
		}
	}
	return y
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
	var r = StepTwo(n)
	// step 3 -- check GCD between 1 and r
	for i := 2; i <= r; i++ {
		gcd, _, _ := GCD(n, i)
		if 1 < gcd && gcd < n {
			return false
		}
	}
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
	//for n := 2; n <= 10000; n++ {
	//	var test = PerfectPower(n)
	//	if test {
	//		fmt.Println(strconv.Itoa(n) + " is a perfect power")
	//	}
	//}
	//fmt.Println(OrderMod(2739, 674893))
	//fmt.Println(stepTwo(29))
	// fmt.Println(EulerTotient(15))
	// test polynomial multiplication
	//x := Polynomial{2, []int{1, 1, 1}}
	//y := Polynomial{2, []int{1, 2, 3}}
	//z := PolynomialMultiply(x, y)
	//for _, i := range z.coefs {
	//	fmt.Println(i)
	//}
}
