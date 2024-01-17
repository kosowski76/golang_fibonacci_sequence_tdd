# golang_fibonacci_sequence_tdd
    Calculating the value of the Fibonacci sequence for the current hour.

  Sample source code calculating the sum of the Fibonacci Sequence for a given hour (by an integer),
presenting the stages of creating 'life cycle' functionality using the Test Driven Development technique.

            0 for n = 0;
    F(n) =  1 for n = 1;
            F(n-2) + F(n-1) for n > 1;


 ## 01.01. Problem formulation and analysis.

 Input data:

    - the input is the function 'getHourOnly()', in the tests it will be necessary 
      to count the hours from 0 to 24 (probably from -1 to 25 also)

    - validation using the function 'checkValidity(int sum, ...)'
        if 'sum' > 0 'println(sum)';
        if 'sum' < 0 'alternativeLogic()';
        return (sum < 0) ? false : true;

 Tests:

    'GivenNegativeHourWhenCountingFiboShouldTrowNFE()'
        if the hour is less than 0 then the function 'FibonacciSequence(number int)'
        throws an error (Exception) - (perform test for n = -1 and other negative numbers)

    'Given0shouldReturn0()' and for other positive n values 'GivenXshouldReturnY()'


 ## 02.01. Designing algorithms and tests.

     The function is supposed to return 1 for hour 1 
    (i.e. we implement hour 0, but we will never return it)

     The function should also be able to count on a specific day of the month, 
    but also be open to further, larger arguments (integration with external API #342)

     The function is intended to handle lack of integrity in the case of a negative number argument


 ## 02.02. Preliminary assurance of test compliance.

     Test preparation (stage Red) and initial assurance
    acceptance of test fulfillment (stage Greenen)
