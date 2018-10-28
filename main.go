package main

import (
	"flag"
	"log"
)

func isPrime(i uint) bool {
	for j := i - 1; j > 1; j-- {
		if i%j == 0 {
			return false
		}
	}
	return true
}

func main() {

	var start uint
	var end uint
	flag.UintVar(&start, "start", 0, "usage")
	flag.UintVar(&end, "end", 10, "usage")

	flag.Parse()
	//declare primes and give it a type
	primes := make([]uint, end)
	for prime := range primes {
		if isPrime(primes[prime]) {
			log.Println(prime)
		}
	}
	log.Println(start)
	log.Println(end)
	for i := start; i < end; i++ {
		if isPrime(i) {
			log.Println(i)
		}
	}
}
