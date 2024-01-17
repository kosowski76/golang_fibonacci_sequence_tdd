package fibonacci_sum // SumMechanics SumSequences
import "errors"

func GetFibonacciSum(n int) (int64, error) {

	if n < 0 {
		return 0, errors.New("invalid argrument: Number Format Exception")
	}

	return getFibonacciInteractiveSum(n), nil
}

/**         0 for n = 0;
 *  F(n) =  1 for n = 1;
 *          F(n-2) + F(n-1) for n > 1;
 **/
func getFibonacciInteractiveSum(n int) int64 {

	if n == 0 {
		return 0
	}

	var twoBehind, oneBehind, current int64 = 0, 0, 1

	for j := 1; j < n; j++ {

		twoBehind = oneBehind
		oneBehind = current
		current = twoBehind + oneBehind
	}

	return current
}

/*
	deprecated

This takes too much time to use, function ready to remove
*/
func getFibonacciRecursive(n int) int64 {

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return getFibonacciRecursive(n-2) + getFibonacciRecursive(n-1)
}
