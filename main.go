package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	"strconv"
)

var count int

func main() {
	if len(os.Args) < 1 {
		panic("not found file to read")
	}
	datafilePath := os.Args[1]

	file, err := os.Open(datafilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := convertToInt64(scanner.Text())
		if err != nil {
			fmt.Println(err.Error())
		} else {
			calculateReducedNumber(number)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func calculateReducedNumber(number int64) {
	count = 1
	showResult := true

	for !isPrimenNumber(number) {
		sum := getSumPrimeFactorNumbers(number)
		if number == sum {
			showResult = false
			break
		}
		number = sum
	}

	if showResult {
		fmt.Println(number, count)
	}
}

func getSumPrimeFactorNumbers(number int64) int64 {
	result := primeFactorization(int(number))

	var sum int64
	for n, q := range result {
		sum += int64(n * q)
	}

	return sum
}

func convertToInt64(text string) (int64, error) {
	n, err := strconv.ParseInt(text, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("%d is not an int64", n)
	}
	return n, nil
}

func isPrimenNumber(number int64) bool {
	return big.NewInt(number).ProbablyPrime(0)
}

func primeFactorization(n int) (pfs map[int]int) {
	count++
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
