package aks

import (
	"math"
)

type EuclidData struct {
	GCD int
	U   int
	V   int
}

// Euclidean algorithm
func EuclidAlgo(a, b int) EuclidData {
	// if b = 0, then the gcd is a
	if b == 0 {
		return EuclidData{a, 1, 0}
	}
	// Keeps tracks of the sign of a and b and makes sure
	// a and b are non-negative
	var na, nb bool
	if a < 0 {
		na = true
		a *= -1
	}
	if b < 0 {
		nb = true
		b *= -1
	}
	// Variables we will return
	var g, u, v int
	u = 1
	g = a
	x := 0 // keeps track of the number a's used in the Euclidean algorithm
	y := b // keeps track of the denominator in the Euclidean algorithm
	for y != 0 {
		t := ModN(uint(y), g) // find t,q with g = qy + t
		q := g / y
		s := u - q*x
		u = x
		g = y
		x = s
		y = t
	}
	v = (g - a*u) / b
	if !na && !nb {
		return EuclidData{g, u, v}
	} else if !na && nb {
		return EuclidData{g, u, -v}
	} else if na && !nb {
		return EuclidData{g, -u, v}
	} else {
		return EuclidData{g, -u, -v}
	}
}

// MillerRabinTest tests a sufficient amount of witnesses such that N may be determined to be composite or prime
// -- assuming the Riemann Hypothesis. false = composite, true = prime
func MillerRabinTest(N int) bool {
	upperbound := 2 * math.Log(float64(N)) * math.Log(float64(N))
	upperbound = math.Min(float64(N), upperbound-1)
	composite := false
	for a := 2; a <= int(upperbound); a++ {
		b := MillerRabinWitness(N, a)
		if b {
			composite = true
			break
		}
	}
	return !composite
}

// MillerRabinWitness tests one witness of the Miller Rabin test.
func MillerRabinWitness(N, a int) bool {
	// Input. Integer n to be tested, integer a as potential
	// witness.
	if N < 0 {
		N *= -1
	}
	// 1. If n is even or 1 < gcd(a,n) < n, return Composite
	if N%2 == 0 {
		return true
	} else if EuclidAlgo(N, a).GCD != 1 {
		return true
	}
	q := N - 1
	k := 0
	// 2. Write n-1 = 2^k q with q odd
	for q%2 != 0 {
		q = q / 2
		k += 1
	}
	// 3. Set a = a^q mod n.
	a = ModN(uint(N), FastPowerMod(uint(N), a, uint(q)))
	// 4. If a = 1 mod n, return Test Fails.
	if a == 1 {
		return false
	}
	// 5. Loop i = 0,1,2,...,k-1
	for i := 0; i < k; i++ {
		//	6. If a = -1 mod n, return Test Fails.
		if (a+1)%N == 0 {
			return false
		}
		//	7. Set a = a^2 mod n
		a = FastPowerMod(uint(N), a, 2)
	}
	// 8. End i loop.
	// 9. Return Composite
	return true
}
