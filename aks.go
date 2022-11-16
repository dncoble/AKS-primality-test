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

// Polynomial struct contains data for representing a polynomial.
// d is the degree of the polynomial
// coefs []int is an integer list of the coefficients, so the list {c_0, c_1, c_2,...,c_d}
// from the polynomial c_0*x^0 + c_1*x^1 + c_2*x^2 + ... + c_d*x^d
type Polynomial struct {
	d     int   //order
	coefs []int // coefficients
}

// PolynomialMultiply multiplies two polynomials represented by Polynomial struct
// timing is product of orders of the polynomials
func PolynomialMultiply(X, Y Polynomial) Polynomial {
	dx := X.d
	dy := Y.d
	coefs := make([]int, dx*dy+1)
	for i := 0; i < len(coefs); i++ {
		coefs[i] = 0
	}
	for i, c1 := range X.coefs {
		for j, c2 := range Y.coefs {
			coefs[i+j] += c1 * c2
		}
	}
	return Polynomial{dx * dy, coefs}
}

// PolynomialAdd does polynomial addition between the two given polynomials
// timing is the minimum of the degrees of the polynomials
func PolynomialAdd(X, Y Polynomial) Polynomial {
	// take the upper limit as the min degree between X and Y
	var poly1 Polynomial
	var poly2 Polynomial
	if X.d < Y.d { // Y has the greater degree
		poly1 = Y
		poly2 = X
	} else {
		poly1 = X
		poly2 = Y
	}
	var coefs = poly1.coefs
	for i := 0; i < poly2.d+1; i++ {
		coefs[i] += poly2.coefs[i]
	}
	return Polynomial{len(coefs) + 1, coefs}
}

// PolynomialMod X mod (Y, N) for X, Y polynomial and N integer, using the polynomial division algorithm
// timing is ___
func PolynomialMod(X, Y Polynomial, N int) Polynomial {
	return Polynomial{1, []int{1}}
}

// PolynomialFastPower X^n mod(Y, N) with X, Y polynomials and n, N integers
// timing is ____
func PolynomialFastPower(X Polynomial, n int, Y Polynomial, N int) Polynomial {
	var b Polynomial
	A := X
	b = Polynomial{0, []int{1}}
	for n > 0 {
		if n%2 == 1 {
			b = PolynomialMod(PolynomialMultiply(b, A), Y, N)
		}
		A = PolynomialMultiply(A, A)
		n = n / 2
	}
	return b
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

func StepFive(n int, r int) bool {
	var upper int = int(math.Floor(math.Sqrt(float64(EulerTotient(r))) * math.Log2(float64(n))))
	// modPolynomial is X^r - 1
	var modPolynomialCoefs = make([]int, r+1)
	for i := 0; i < r+1; i++ { // instantiate modPolynomial coefs
		if i == 0 {
			modPolynomialCoefs[i] = -1
		} else if i == r {
			modPolynomialCoefs[i] = 1
		} else {
			modPolynomialCoefs[i] = 0
		}
	}
	var modPolynomial = Polynomial{r, modPolynomialCoefs}
	for a := 1; a <= upper; a++ {
		// left polynomial is (X+a)^n mod
		var leftPolynomial = PolynomialFastPower(Polynomial{1, []int{a, 1}}, n, modPolynomial, n)
		// right polynomial is X^n + a
		var rightPolynomial = PolynomialFastPower(Polynomial{1, []int{0, 1}}, n, modPolynomial, n)
		rightPolynomial = PolynomialAdd(rightPolynomial, Polynomial{0, []int{a}})
		if !PolynomialEquality(leftPolynomial, rightPolynomial) {
			return false
		}
	}
	return true
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
	return StepFive(n, r)
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
