- always seperate the domain code from the sideeffects code
- named return values. these will created in a function scope and will set to the default value

## Rules for writing test cases 


- It needs to be in a file name xxx_test.go
- The test should start with word `Test`
- The test should take only one argument of type `*testing.T`


## Discipline

Let's go over the cycle again

- Write a test
- Make the compiler pass
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass
- Refactor


- Public function should always be started with capital letter
- Private function should always be started with small letter