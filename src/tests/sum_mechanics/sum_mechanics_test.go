package sum_mechanics_test

import "github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sequence"

// todo: test if another negative number works
func GivenNegativeHourWhenCountingFiboShouldTrowNFE() {

	number := -1
	// calculateFor(number)
	fibonacci_sequence.FibonacciSequence(number)
	// todo: should throw NumberFormatException
}

func Given0shouldReturn0() {
}
func Given1shouldReturn1() {
}
func Given2shouldReturn1() {
}
func Given3shouldReturn2() {
}
func Given10shouldReturn54() {
}
func Given9shouldReturn34() {
}
