package sum_mechanics_test

import (
	"github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sequence"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: test if another negative number works
func TestGivenNegativeHourWhenCountingFiboShouldTrowNFE(t *testing.T) {

	number := -1
	// calculateFor(number)
	fibonacci_sequence.FibonacciSequence(number)
	// todo: should throw NumberFormatException
}

func TestGiven0shouldReturn0(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(0)
	assert.EqualValues(t, 0, fibonacciCalc)
}
func TestGiven1shouldReturn1(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(1)
	assert.EqualValues(t, 1, fibonacciCalc)
}
func TestGiven2shouldReturn1(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(2)
	assert.EqualValues(t, 1, fibonacciCalc)
}
func TestGiven3shouldReturn2(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(3)
	assert.EqualValues(t, 2, fibonacciCalc)
}
func TestGiven10shouldReturn54(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(10)
	assert.EqualValues(t, 54, fibonacciCalc)
}
func TestGiven9shouldReturn34(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(9)
	assert.EqualValues(t, 34, fibonacciCalc)
}
func TestGiven23shouldReturn28657(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(23)
	assert.EqualValues(t, 28657, fibonacciCalc)
}
func TestGiven24shouldReturn46368(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(24)
	assert.EqualValues(t, 46368, fibonacciCalc)
}
func TestGiven25shouldReturn75025(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(25)
	assert.EqualValues(t, 75025, fibonacciCalc)
}
func TestGiven27shouldReturn196418(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(27)
	assert.EqualValues(t, 196418, fibonacciCalc)
}
func TestGiven28shouldReturn317811(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(28)
	assert.EqualValues(t, 317811, fibonacciCalc)
}
func TestGiven29shouldReturn514229(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(29)
	assert.EqualValues(t, 514229, fibonacciCalc)
}
func TestGiven30shouldReturn832040(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(30)
	assert.EqualValues(t, 832040, fibonacciCalc)
}
func TestGiven31shouldReturn1346269(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(31)
	assert.EqualValues(t, 1346269, fibonacciCalc)
}
func TestGiven47shouldReturn2971215073(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(47)
	assert.EqualValues(t, 2971215073, fibonacciCalc)
}
func TestGivenNegative1shouldReturnException(t *testing.T) {
	fibonacciCalc := fibonacci_sequence.FibonacciSequence(-1)
	assert.EqualValues(t, 0000, fibonacciCalc)
}
