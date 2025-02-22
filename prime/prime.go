package main

import "fmt"

func main() {
	fmt.Println(prime(5))
	fmt.Println(prime(4))
}

func prime(n int) string {
	isPrime := true

	if n < 2 {
		isPrime = false
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			isPrime = false
			break
		}
	}

	if isPrime {
		return fmt.Sprintf("%d is prime", n)
	} else {
		return fmt.Sprintf("%d is not prime", n)
	}
}
