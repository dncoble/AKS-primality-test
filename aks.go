func PerfectPower(n int) bool {

	return false
}
// AKS algorithm false = composite, true = prime
func AKS(n int) bool {
	// step 1 -- find if n is a perfect power
	if(PerfectPower(n)) {
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