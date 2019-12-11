package main

import "testing"

func TestEcho(t *testing.T) {
	records := [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"7","8","9"},
	}
	expected := "1,2,3\n4,5,6\n7,8,9"
	result := Echo(records)
    if result != expected {
       t.Errorf("Echo was incorrect, got: %s, expected: %s.", result, expected)
	}

	records = [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"X","Y","Z"},
	}
	expected = "1,2,3\n4,5,6\nX,Y,Z"
	result = Echo(records)
    if result != expected {
       t.Errorf("Echo was incorrect, got: %s, expected: %s.", result, expected)
    }
}

func TestInvert(t *testing.T) {
	records := [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"7","8","9"},
	}
	expected := "1,4,7\n2,5,8\n3,6,9"
	result := Invert(records)
    if result != expected {
       t.Errorf("Invert was incorrect, got: %s, expected: %s.", result, expected)
	}

	records = [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"X","Y","Z"},
	}
	expected = "1,4,X\n2,5,Y\n3,6,Z"
	result = Invert(records)
    if result != expected {
       t.Errorf("Invert was incorrect, got: %s, expected: %s.", result, expected)
    }
}

func TestFlatten(t *testing.T) {
	records := [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"7","8","9"},
	}
	expected := "1,2,3,4,5,6,7,8,9"
	result := Flatten(records)
    if result != expected {
       t.Errorf("Flatten was incorrect, got: %s, expected: %s.", result, expected)
	}

	records = [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"X","Y","Z"},
	}
	expected = "1,2,3,4,5,6,X,Y,Z"
	result = Flatten(records)
    if result != expected {
       t.Errorf("Flatten was incorrect, got: %s, expected: %s.", result, expected)
    }
}

func TestSum(t *testing.T) {
	records := [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"7","8","9"},
	}
	expected := "45"
	result, err := Sum(records)
    if result != expected {
       t.Errorf("Sum was incorrect, got: %s, expected: %s.", result, expected)
	}
	if err != nil {
		t.Errorf("non-nil Error returned: %s", err.Error())
	}

	records = [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"X","Y","Z"},
	}
	expected = ""
	result, err = Sum(records)
    if result != expected {
       t.Errorf("Sum was incorrect, got: %s, expected: %s.", result, expected)
    }
	if err == nil {
		t.Errorf("nil Error returned, expected non-nil")
	}
}

func TestMultiply(t *testing.T) {
	records := [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"7","8","9"},
	}
	expected := "362880"
	result, err := Multiply(records)
    if result != expected {
       t.Errorf("Multiply was incorrect, got: %s, expected: %s.", result, expected)
	}
	if err != nil {
		t.Errorf("non-nil Error returned: %s", err.Error())
	}

	records = [][]string{
		{"1","2","3"},
		{"4","5","6"},
		{"X","Y","Z"},
	}
	expected = ""
	result, err = Multiply(records)
    if result != expected {
       t.Errorf("Multiply was incorrect, got: %s, expected: %s.", result, expected)
    }
	if err == nil {
		t.Errorf("nil Error returned, expected non-nil")
	}
}
