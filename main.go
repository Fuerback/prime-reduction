package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./sample.in")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, _ := convertToInt(scanner.Text())
		itIsPrimenNumber(number)
		result := primeFactorization(int(number))
		for k, _ := range result {
			fmt.Println(k)
			// somar os numeros do mapa e conferir se o total é primo, caso nao seja, passar pelo refactor e repetir o processo.. caso seja primo, printar o numero e o numero de vezes que passou pelo refactor
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func convertToInt(text string) (int64, error) {
	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%d is not an int64", n)
	}
	return n, nil
}

func itIsPrimenNumber(number int64) bool {
	if big.NewInt(number).ProbablyPrime(0) {
		fmt.Println(number, "is prime")
	} else {
		fmt.Println(number, "is not prime")
	}
	return false
}

func primeFactorization(n int) (pfs map[int]int) {
	pfs = make(map[int]int)

	// Get the number of 2s that divide n
	for n%2 == 0 {
		if _, ok := pfs[2]; ok {
			pfs[2] += 1
		} else {
			pfs[2] = 1
		}
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			if _, ok := pfs[i]; ok {
				pfs[i] += 1
			} else {
				pfs[i] = 1
			}
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs[n] = 1
	}

	return
}
