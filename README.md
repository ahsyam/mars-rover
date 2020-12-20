#Mars-rover problem

Implementation for mars-rover with TDD

## implementation notes:
I decided to use  Go testing & gin microframework and go mock (if needed)
based on recommendation in the Email, I tried to focus on simplicity rather than complex solution

## solution steps
1. I decided that I'll start with solution for the problem first by writing the tests and minimal code to let them pass.
2. I'll write the implementation.
3. Write the tests for the API and minimal code to pass the tests.
4. Implement the endpoints and do the refactor to the code.
5. Will start with part 2 and part 3 the same way

## how to run the code
application will be dockerized and i added some make commands to simplify things.

* to build the project `make build`
* to run the project `make run`
* to clean the build `make clean`
* swagger for the API `make run-swagger`
* to run the shell `make shell`
* to see the logs `make watch-logs`

## How to test the API
you can use swagger to do that