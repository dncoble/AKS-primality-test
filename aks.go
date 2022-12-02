package main

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

// FastPower perform the without modulo by N)
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

// FastPowerMod performs the fast power algorithm modulo N
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

// GCD of a, b using the Euclidean algorithm, returns (gcd, u, v) where u v are such that u*a + v*b = (gcd)
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

// OrderMod returns the order of a modulo r. The order only exists when gcd(a, r) = 1, so if that is not true,
// the function returns 0
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

// EulerTotient function finds the number of numbers less than x that are relatively prime to x
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

// Polynomial struct contains data for representing a polynomial.
// d is the degree of the polynomial
// coefs []int is an integer list of the coefficients, so the list {c_d, c_{d-1},,...,c_0}
// from the polynomial c_0*x^0 + c_1*x^1 + c_2*x^2 + ... + c_d*x^d
type Polynomial struct {
	d     int   //order
	coefs []int // coefficients
}

// PolynomialEquality check whether degree and all coefficients of two polynomials are equal
func PolynomialEquality(X, Y Polynomial) bool {
	if X.d != Y.d {
		return false
	}
	for i := 0; i < X.d+1; i++ {
		if X.coefs[i] != Y.coefs[i] {
			return false
		}
	}
	return true
}

// PolynomialMultiply multiplies two polynomials represented by Polynomial struct
func PolynomialMultiply(X, Y Polynomial) Polynomial {
	dx := X.d
	dy := Y.d
	coefs := make([]int, dx+dy+1)
	for i := 0; i < len(coefs); i++ {
		coefs[i] = 0
	}
	for i, c1 := range X.coefs {
		for j, c2 := range Y.coefs {
			coefs[i+j] += c1 * c2
		}
	}
	return Polynomial{len(coefs) - 1, coefs}
}

// PolynomialAdd does polynomial addition between the two given polynomials
func PolynomialAdd(X, Y Polynomial) Polynomial {
	// take the upper limit as the min degree between X and Y
	var poly1 Polynomial
	var poly2 Polynomial
	var d int
	if X.d < Y.d { // Y has the greater degree
		poly1 = Y
		poly2 = X
		d = Y.d - X.d
	} else { // X has the greater degree
		poly1 = X
		poly2 = Y
		d = X.d - Y.d
	}
	var coefs = poly1.coefs
	for i := 0; i < poly2.d+1; i++ {
		coefs[i+d] += poly2.coefs[i]
	}
	// remove leading zeros from polynomial coefficients
	var j = 0
	for coefs[j] == 0 {
		j++
	}
	coefs = coefs[j:]
	return Polynomial{len(coefs) - 1, coefs}
}

// PolynomialMod X mod (Y, N) for X, Y polynomial and N integer, using the polynomial division algorithm
func PolynomialMod(X, Y Polynomial, N int) Polynomial {
	// calculate X mod Y using PolynomialRemainder
	m := PolynomialRemainder(X, Y)
	// take all coefficients modulo N
	for c := 0; c < len(m.coefs); c++ {
		m.coefs[c] = ModN(uint(N), m.coefs[c])
	}
	// remove leading zeros from polynomial coefficients
	coefs := m.coefs
	var j = 0
	for coefs[j] == 0 {
		j++
	}
	coefs = coefs[j:]
	return Polynomial{len(coefs) - 1, coefs}
}

// PolynomialFastPower X^n mod(Y, N) with X, Y polynomials and n, N integers
func PolynomialFastPower(X Polynomial, n int, Y Polynomial, N int) Polynomial {
	var b Polynomial
	A := X
	b = Polynomial{0, []int{1}}
	for n > 0 {
		if n%2 == 1 {
			b = PolynomialMod(PolynomialMultiply(b, A), Y, N)
		}
		A = PolynomialMod(PolynomialMultiply(A, A), Y, N)
		n = n / 2
	}
	return b
}

// PolynomialRemainder returns the remainder after Polynomial i is divided by Polynomial N
// using the polynomial long division algorithm
func PolynomialRemainder(i Polynomial, N Polynomial) Polynomial {
	orderi := i.d
	orderN := N.d
	// if the order of i is already less than N, return i
	if orderi < orderN {
		return i
	}
	ci := i.coefs
	cN := N.coefs
	temp := i
	tempCoef := make([]int, len(ci), len(ci))
	for i := 0; i < len(cN); i++ {
		tempCoef[i] = cN[i]
	}
	var scalar int
	currentOrder := orderi - orderN
	for currentOrder >= 0 {
		scalar = int((math.Floor(float64(temp.coefs[0] / cN[0])))) * (-1)
		for i := range cN {
			tempCoef[i] = cN[i] * scalar
		}
		if len(tempCoef) != len(temp.coefs) {
			tempCoef = tempCoef[:len(temp.coefs)]
		}
		temp = PolynomialAdd(temp, Polynomial{currentOrder + orderN, tempCoef})
		currentOrder = temp.d - orderN
	}
	return temp
}

// PerfectPower determines whether n can be represented as a perfect power a^b
// used in step 1 of the AKS algorithm
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

// StepTwo of the AKS algorithm
func StepTwo(n int) int {
	var lower = int(math.Ceil(math.Log2(float64(n)) * math.Log2(float64(n))))
	var r = 2
	for {
		if OrderMod(n, r) > lower {
			return r
		}
		r++
	}
}

// StepFive of the AKS algorithm
func StepFive(n int, r int) bool {
	var upper int = int(math.Floor(math.Sqrt(float64(EulerTotient(r))) * math.Log2(float64(n))))
	// modPolynomial is X^r - 1
	var modPolynomialCoefs = make([]int, r+1)
	modPolynomialCoefs[0] = 1
	modPolynomialCoefs[r] = -1
	var modPolynomial = Polynomial{r, modPolynomialCoefs}
	for a := 1; a <= upper; a++ {
		// left polynomial is (X+a)^n mod
		var leftPolynomial = PolynomialFastPower(Polynomial{1, []int{1, a}}, n, modPolynomial, n)
		// right polynomial is X^n + a
		var rightPolynomial = PolynomialFastPower(Polynomial{1, []int{1, 0}}, n, modPolynomial, n)
		rightPolynomial = PolynomialAdd(rightPolynomial, Polynomial{0, []int{a}})
		if !PolynomialEquality(leftPolynomial, rightPolynomial) {
			return false
		}
	}
	return true
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
	return StepFive(n, r)
}
