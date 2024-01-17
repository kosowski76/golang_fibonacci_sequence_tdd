package main

import (
	"fmt"
	"github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sequence"
	"time"
)

var (
	hourOnly           int
	stringElementValue int
)

func main() {

	hourOnly = getHourOnly()
	fmt.Println("Current hour golden Fibonacci for n =", hourOnly)

	stringElementValue = fibonacci_sequence.FibonacciSequence(hourOnly)

	fmt.Println(stringElementValue)
}

/**
 * Return the local system time of day in 24-hour format
 **/
func getHourOnly() int {

	now := time.Now()

	hour := now.Hour()

	return hour
}
