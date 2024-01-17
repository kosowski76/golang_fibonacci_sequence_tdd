package fibonacci_sequence // SumMechanics SumSequences

func FibonacciSequence(hour int) int {

	if hour == 0 {
		return 0
	}
	if hour == 1 || hour == 2 {
		return 1
	}

	return FibonacciSequence(hour-2) + FibonacciSequence(hour-1)
}

/**         0 for n = 0;
 *  F(n) =  1 for n = 1;
 *          F(n-2) + F(n-1) for n > 1;
 **/
