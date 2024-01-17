package sum_mechanics_test

import (
	"github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

// todo: test if another negative number works
func TestGivenNegativeHourWhenCountingFiboShouldTrowNFE(t *testing.T) {

	number := -1
	// calculateFor(number)
	_, error := fibonacci_sum.GetFibonacciSum(number)
	// todo: should throw NumberFormatException
	error.Error()
}

func TestGiven0shouldReturn0(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(0)
	assert.EqualValues(t, 0, fibonacciCalc)
}
func TestGiven1shouldReturn1(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(1)
	assert.EqualValues(t, 1, fibonacciCalc)
}
func TestGiven2shouldReturn1(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(2)
	assert.EqualValues(t, 1, fibonacciCalc)
}
func TestGiven3shouldReturn2(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(3)
	assert.EqualValues(t, 2, fibonacciCalc)
}
func TestGiven10shouldReturn54(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(10)
	assert.EqualValues(t, 54, fibonacciCalc)
}
func TestGiven9shouldReturn34(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(9)
	assert.EqualValues(t, 34, fibonacciCalc)
}
func TestGiven23shouldReturn28657(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(23)
	assert.EqualValues(t, 28657, fibonacciCalc)
}
func TestGiven24shouldReturn46368(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(24)
	assert.EqualValues(t, 46368, fibonacciCalc)
}
func TestGiven25shouldReturn75025(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(25)
	assert.EqualValues(t, 75025, fibonacciCalc)
}
func TestGiven27shouldReturn196418(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(27)
	assert.EqualValues(t, 196418, fibonacciCalc)
}
func TestGiven28shouldReturn317811(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(28)
	assert.EqualValues(t, 317811, fibonacciCalc)
}
func TestGiven29shouldReturn514229(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(29)
	assert.EqualValues(t, 514229, fibonacciCalc)
}
func TestGiven30shouldReturn832040(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(30)
	assert.EqualValues(t, 832040, fibonacciCalc)
}
func TestGiven31shouldReturn1346269(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(31)
	assert.EqualValues(t, 1346269, fibonacciCalc)
}
func TestGiven47shouldReturn2971215073(t *testing.T) {
	fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(47)
	assert.EqualValues(t, 2971215073, fibonacciCalc)
}
