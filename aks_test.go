package aks

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func test(t *testing.T) {
	agreeAllNumbers := true
	start := time.Now()
	// test first 10,000 numbers
	for i := 2; i <= 10000; i++ {
		AKSResult := AKS(i)
		MillerRabinResult := MillerRabinTest(i)
		if AKSResult == MillerRabinResult {
			//if AKSResult {
			//	fmt.Println("AKS and Miller Rabin agree " + strconv.Itoa(i) + " is prime.")
			//} else {
			//	fmt.Println("AKS and Miller Rabin agree " + strconv.Itoa(i) + " is not prime.")
			//}
		} else {
			fmt.Println("AKS and Miller Rabin disagree for " + strconv.Itoa(i))
			agreeAllNumbers = false
		}
	}
	if agreeAllNumbers {
		fmt.Println("AKS and Miller Rabin tests agree for integers up to 10,000")
	}
	timeElapsed := time.Since(start)
	fmt.Printf("Time elapsed: %s\n", timeElapsed)
}
