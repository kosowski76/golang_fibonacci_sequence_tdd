package fibonacci_sum // SumMechanics SumSequences
import "errors"

func GetFibonacciSum(n int) (int64, error) {

	if n < 0 {
		return 0, errors.New("invalid argrument: Number Format Exception")
	}

	return getFibonacciRecursive(n), nil
}

/**         0 for n = 0;
 *  F(n) =  1 for n = 1;
 *          F(n-2) + F(n-1) for n > 1;
 **/
func getFibonacciRecursive(n int) int64 {

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return getFibonacciRecursive(n-2) + getFibonacciRecursive(n-1)
}

/*
	deprecated

ready to remove
*/
func FibonacciSequence(n int) int {

	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}

	return FibonacciSequence(n-2) + FibonacciSequence(n-1)
}
