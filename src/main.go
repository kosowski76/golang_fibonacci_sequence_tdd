package main

import (
	"fmt"
	"github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sum"
	"time"
)

var (
	hourOnly           int
	stringElementValue int64
)

func main() {

	hourOnly = getHourOnly()
	fmt.Println("Current hour golden Fibonacci for n =", hourOnly)

	stringElementValue, _ = fibonacci_sum.GetFibonacciSum(hourOnly)

	fmt.Println(stringElementValue)
}

/**
 * Return the local system time of day in 24-hour format
 **/
func getHourOnly() int {

	now := time.Now()
	return now.Hour()
}
