# AKS primality test
The [AKS primality test](https://www.cse.iitk.ac.in/users/manindra/algebra/primality_v6.pdf) (discovered by Agrawal, Kayal, and Saxena) is the first discovered polynomial-time algorithm for testing primality of a given number.

## Testing the code
`main` in AKS.go runs checks primality of the first 10,000 natural numbers. To check whether the AKS algorithm produces the correct answer, primality is also tested with the Miller-Rabin test. The Miller-Rabin test is assured to produce the correct answer so long as the Riemann Hypothesis is true. 

## How it works

Input: integer n > 1.

1. If ( $n = a^b$ for $a \in \mathcal{N}$ and $b>1$), output `COMPOSITE`.
2. Find the smallest $r$ such that $o_r(n) > \log^2 n$.
3. If $1 < (a, n) < n$ for some $a \leq r$, output `COMPOSITE`.
4. If $n \leq r$ output `PRIME`.
5. For $a = 1$ to $\lfloor \sqrt{\phi(r)}\log n\rfloor$ do
    
    if ( $(X+a)^n \neq X^n+a \quad (\mathrm{mod} X^r-1, n)$ ), output `COMPOSITE`.
6. Output `PRIME`.

## Functions and timing analysis

Computational complexity of Polynomial functions is given in terms of the order of the polynomial, while other functions are given in the size of the integer input. All exponential functions are called by `AKS` with $\log$-bounded inputs, producing in total a polynomial-time algorithm.
|Function|Description|Timing|
|--------|-----------|------|
|`ModN`| $m$ such that $0 \leq m < N$ and $i = m \mod N$ | $\mathcal O (1)$ |
|`FastPower`| fast powering algorithm | $\mathcal O (\log(n))$ |
|`FastPowerMod`| fast powering algorithm with modulus | $\mathcal O (\log(n))$ |
|`GCD`| greatest common divisor of two numbers | $\mathcal O (\log(n))$ |
|`OrderMod`| the order of a modulo r | $\mathcal O (n\log(n))$ |
|`EulerTotient`| number of numbers less than x relatively prime to x | $\mathcal O (n\log(n))$ |
|`PolynomialEquality`| check equality between polynomials | $\mathcal O (n)$ |
|`PolynomialAdd`| addition on polynomial objects | $\mathcal O (n)$|
|`PolynomialMultiply`| multiplication on polynomial objects | $\mathcal O (n^2)$ |
|`PolynomialMod`| $X \mod Y, N$ for $X$, $Y$ polynomial and $N$ integer | $\mathcal O (n^2)$ |
|`PolynomialFastPower`| fast powering algorithm on polynomial objects | $\mathcal O (n^3)$ |
|`PolynomialRemainder`| remainder after dividing two polynomials | $\mathcal O (n^2)$ |
|`StepTwo`| Step 2 of the AKS algorithm | $\mathcal O (\log^5(n))$* |
|`StepFive`| Step 5 of the AKS algorithm | $\mathcal O (\log^{15/2}(n))$**|
|`PerfectPower`| tests if $n$ is of the form $a^b$ | $\mathcal O (\log^3(n))$ |
|`AKS`| AKS primality test | $\mathcal O (\log^{15/2}(n))$**|

\* $\mathcal O (\log^2(n))$ if Artin's conjecture is proven true.

** $\mathcal O (\log^{6}(n))$ if Artin's conjecture is proven true


## Citations
[PRIMES is in P. <em> Agrawal, Kayal, and Saxena. </em>  Annals of Mathematics, pg. 781-793. Volume 160, Issue 2.](https://www.cse.iitk.ac.in/users/manindra/algebra/primality_v6.pdf)

## License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
