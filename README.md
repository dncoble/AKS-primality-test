# AKS primality test
The [AKS primality test](https://www.cse.iitk.ac.in/users/manindra/algebra/primality_v6.pdf) (discovered by Agrawal, Kayal, and Saxena) is the first discovered polynomial-time algorithm for testing primality of a given number.

## Current progress
A draft code has been written which goes through all five steps, but it has not been checked/debugged.

## Testing the code
main.go runs a program which checks primality of the first 10,000 natural numbers. To check whether the AKS algorithm produces the correct answer, primality is also tested with the Miller-Rabin test. The Miller-Rabin test is assured to produce the correct answer so long as the Riemann Hypothesis is true. 

## How it works

Input: integer n > 1.

1. If ( $n = a^b$ for $a \in \mathcal{N}$ and $b>1$), output `COMPOSITE`.
2. Find the smallest $r$ such that $o_r(n) > \log^2 n$.
3. If $1 < (a, n) < n$ for some $a \leq r$, output `COMPOSITE`.
4. If $n \leq r$ output `PRIME`.
5. For $a = 1$ to $\lfloor \sqrt{\phi(r)}\log n\rfloor$ do
    
    if ( $(X+a)^n \neq X^n+a \quad (\operatorname{mod} X^r-1, n)$ ), output `COMPOSITE`.
6. Output `PRIME`.

## Functions and timing analysis

Computational complexity of Polynomial functions is given in terms of the order of the polynomial, while other functions are given in the size of the integer input.
|Function|Description|Timing|
|--------|-----------|------|
|`ModN`| | $\mathcal O ()$ |
|`GCD`| | $\mathcal O ()$ |
|`OrderMod`| | $\mathcal O ()$ |
|`PolynomialMultiply`| | $\mathcal O (n^2)$ |
|`PolynomialAdd`| | $\mathcal O (n)$|
|`PolynomialMod`| | $\mathcal O ()$ |
|`PolynomialFastPower`| | $\mathcal O ()$ |
|`PolynomialEquality`| | $\mathcal O (n)$ |
|`PolynomialRemainder`| | $\mathcal O ()$ |
|`FastPower`| | $\mathcal O (\log(n))$ |
|`FastPowerMod`| | $\mathcal O ()$ |
|`StepTwo`| | $\mathcal O ()$ |
|`StepFive`| | $\mathcal O ()$ |
|`EulerTotient`| | $\mathcal O ()$ |
|`PerfectPower`| | $\mathcal O ()$ |
|`AKS`| | $\mathcal O ()$ |



## Citations

## License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
