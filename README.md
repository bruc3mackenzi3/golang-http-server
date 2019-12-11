
# League Backend Challenge

## Solution
`main.go` contains the HTTP server and of course main function.  The HTTP Handlers are exposed publicly to allow for unit testing (more on that below).

`matrix.go` contains the matrix API implementing the five matrix operations.  It provides a clean, consistent interface used by the Server.  It is unit tested in `matrix_test.go`.

Additional CSV files have been provided to illustrate functionality and provide
the ability to do more thorough functional testing manually.

To keep things simple a flat directory structure is used here.  By keeping all modules in the same package having to import is avoided.  More substantial Go apps should use the standard workspace directory structure.

## Run Instructions
Download and extract the app and cd into it.  Run the app with:
```
go build
./go.exe
```
Once the web server is up and running open another terminal, cd to the same directory and test with the following curl command:
` curl -F 'file=@matrix.csv' 'http://localhost:8080/echo'`

## Testing
Unit tests have been implemented using Go's `testing` package.
Run the unit tests with the command:
`go test`

### Coverage
The matrix operations have complete code coverage with both expected and erroneous cases tested.  See `matrix_test.go`

An attempt as unit testing the HTTP Handlers was made but without `multipart/form-data` being properly set by the test code it isn't functioning.  Because the CSV parser checks this request header field it returns an error that is isn't set.

The test module `handler_test.go` was included to illustrate the use of the `testing` and `httptest` modules combined.  It also serves as an example of the challenges unit testing modules that are tightly coupled with other APIs or services.  This is typically a scenario best tested with a broader testing methodology such as functional testing or black box testing.

To produce a more thorough test plan an additional file `matrix_invalid.csv` is included.  This aids in manual testing by covering additional cases to those with `matrix.csv`


# Original Problem
In main.go you will find a basic web server written in GoLang. It accepts a single request _/echo_. Extend the webservice with the ability to perform the following operations

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
- Return the matrix as a string in matrix format.
```
// Expected output
1,2,3
4,5,6
7,8,9
```

2. Invert
- Return the matrix as a string in matrix format where the columns and rows are inverted
```
// Expected output
1,4,7
2,5,8
3,6,9
```

3. Flatten
- Return the matrix as a 1 line string, with values separated by commas.
```
// Expected output
1,2,3,4,5,6,7,8,9
```

4. Sum
- Return the sum of the integers in the matrix
```
// Expected output
45
```

5. Multiply
- Return the product of the integers in the matrix
```
// Expected output
362880
```

The input file to these functions is a matrix, of any dimension where the number of rows are equal to the number of columns (square). Each value is an integer, and there is no header row. matrix.csv is example valid input.

Run web server
```
go run .
```

Send request
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## What we're looking for
- The solution runs
- The solution performs all cases correctly
- The code is easy to read
- The code is reasonably documented
- The code is tested
- The code is robust and handles invalid input and provides helpful error messages