# Prime number reduction

A prime number p≥2 is an integer which is evenly divisible by only two integers: 1 and p. A composite integer is one which is not prime. The fundamental theorem of arithmetic says that any integer x can be expressed uniquely as a set of prime factors – those prime numbers which, when multiplied together, give x. Consider the prime factorization bellow:

```sh
231 = 3 × 7 × 11
```

The program reads a file entry, verifying that the line is a number and for each number are summed reduced prime numbers until the sum is not a prime number.

1. if `x` is prime, print x and stop
2. factor `x` into its prime factors `p1,p2,…,pk`
3. let `x`=`p1+p2+⋯+pk`
4. go back to step 1

**The result will be printed followed by the number of times the first line of the process executed.**

# Usage

```sh
go run .\main.go .\sample.in
```

# Output example

## Input

    2
    3
    5
    76
    100
    2001
    4

## Output

    2 1
    3 1
    5 1
    23 2
    5 5
    5 6