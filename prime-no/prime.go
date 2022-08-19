package main

import "fmt"

func main() {
	primeArr := sieveOfEratosthenes(50)
	n := 5
	fmt.Printf("\nn'th prime number is = %v", primeArr[n-1])
}

func sieveOfEratosthenes(n int) []int {
	pno := []int{}
	prime := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		prime[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if prime[p] {
			for i := p * p; i <= n; i += p {
				prime[i] = false
			}
		}
	}

	//print all prime number
	for p := 2; p <= n; p++ {
		if prime[p] {
			fmt.Printf("%v ", p)
			pno = append(pno, p)
		}
	}
	return pno
}
