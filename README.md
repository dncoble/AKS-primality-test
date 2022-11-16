# AKS primality test
The [AKS primality test](https://www.cse.iitk.ac.in/users/manindra/algebra/primality_v6.pdf) (discovered by Agrawal, Kayal, and Saxena) is the first discovered polynomial-time algorithm for testing primality of a given number.

## How to use

## How it works

Input: integer n > 1.

1. If ( $n = a^b$ for $a \in \mathcal{N}$ and $b>1$), output `COMPOSITE`.
2. Find the smallest $r$ such that $o_r(n) > \log^2 n$.
3. If $1 < (a, n) < n$ for some $a \leq r$, output `COMPOSITE`.
4. If $n \leq r$ output `PRIME`.
5. For $a = 1$ to $\lfloor \sqrt{\phi(r)}\log n\rfloor$ do
    
    if ( $(X+a)^n \neq X^n+a \quad (\operatorname{mod} X^r-1, n)$ ), output `COMPOSITE`.
6. Output `PRIME`.


## Citations

## License
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
