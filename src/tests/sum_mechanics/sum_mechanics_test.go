package sum_mechanics_test

import (
	"fmt"
	"github.com/kosowski76/golang_fibonacci_sequence_tdd/src/domain/fibonacci_sum"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGivenNegativeHoursThenFibonacciSumShouldTrowNFE(t *testing.T) {

	var testsNegativeNumbers = []struct {
		n int
	}{
		{-1},
		{-2},
		{-5},
	}

	for _, negativeNum := range testsNegativeNumbers {

		testname := fmt.Sprintf("%d", negativeNum.n)
		t.Run(testname, func(t *testing.T) {

			_, error := fibonacci_sum.GetFibonacciSum(negativeNum.n)
			// todo: should throw NumberFormatException
			error.Error()
		})
	}
}

func TestGivenNShouldReturnY(t *testing.T) {

	var testsPositiveNumbers = []struct {
		n      int
		expect int64
	}{
		{0, 0},
		{1, 1},
		{2, 1},
		{3, 2},
		{4, 3},
		{5, 5},
		{6, 8},
		{9, 34},
		{10, 55},
		{23, 28657},
		{24, 46368},
		{25, 75025},
		{26, 121393},
		{27, 196418},
		{28, 317811},
		{29, 514229},
		{30, 832040},
		{31, 1346269},
		{32, 2178309},
		{35, 9227465},
		{47, 2971215073},
		//		{48, 4807526976},
		//		{49, 7778742049},
		//		{55, 139583862445},
		//		{66, 27777890035288},
	}

	for _, positiveNum := range testsPositiveNumbers {

		testname := fmt.Sprintf("%d", positiveNum.n)
		t.Run(testname, func(t *testing.T) {

			fibonacciCalc, _ := fibonacci_sum.GetFibonacciSum(positiveNum.n)
			assert.EqualValues(t, positiveNum.expect, fibonacciCalc)
		})
	}
}
