package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
The matrix module provides and API for performing various operations on a
matrix.

For each function the input is a 2D array of strings.  The output is a
serialized string representation of the transformed matrix.
*/

func Echo(records [][]string) string {
	var response string
	for _, row := range records {
		response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
	}
	return response[:len(response)-1]  // remove trailing newline
}

func Invert(records [][]string) string {
	// First initialize inverted array
	nRow, nCol := len(records), len(records[0])
	result := make([][]string, nCol)
	for i := 0; i < nCol; i++ {
		result[i] = make([]string, nRow)
	}

	// next copy the values inverting along the way
	for r, row := range records {
		for c, value := range row {
			result[c][r] = value
		}
	}

	// Finally use the echo function to convert 2D array to string
	return Echo(result)
}

func Flatten(records [][]string) string {
	var response string
	for _, row := range records {
		for _, value := range row {
			response += value + ","
		}
	}
	return response[:len(response)-1]  // remove trailing comma
}

/*
Given the 2D array records parse each value to an int, sum the total, and return
as a string.
An error is returned if an invalid int is encountered.
*/
func Sum(records [][]string) (string, error) {
	var result int = 0
	for _, row := range records {
		for _, value := range row {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return "", err
			}
			result += valueInt
		}
	}
	return strconv.Itoa(result), nil
}

/*
Given the 2D array records parse each value to an int, multiply the total, and\
return as a string.
An error is returned if an invalid int is encountered.
*/
func Multiply(records [][]string) (string, error) {
	var result int = 1
	for _, row := range records {
		for _, value := range row {
			valueInt, err := strconv.Atoi(value)
			if err != nil {
				return "", err
			}
			result *= valueInt
		}
	}
	return strconv.Itoa(result), nil
}
