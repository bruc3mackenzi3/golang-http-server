package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

// 1. ECHO
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error reading CSV: %s", err.Error())))
		return
	}

	response := Echo(records)
	fmt.Fprint(w, "\n"+response)
}

// 2. INVERT
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error reading CSV: %s", err.Error())))
		return
	}

	response := Invert(records)
	fmt.Fprint(w, "\n"+response)
}

// 3. FLATTEN
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error reading CSV: %s", err.Error())))
		return
	}

	response := Flatten(records)
	fmt.Fprint(w, "\n"+response)
}

// 4. SUM
func SumHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error reading CSV: %s", err.Error())))
		return
	}

	response, err := Sum(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error computing sum: %s", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response)
	fmt.Println("DEBUG", r.URL.Path, response)
}

// 5. MULTIPLY
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error reading CSV: %s", err.Error())))
		return
	}

	response, err := Multiply(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error computing multiply: %s", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response)
}

func main() {
	http.HandleFunc("/echo", EchoHandler)
	http.HandleFunc("/invert", InvertHandler)
	http.HandleFunc("/flatten", FlattenHandler)
	http.HandleFunc("/sum", SumHandler)
	http.HandleFunc("/multiply", MultiplyHandler)

	fmt.Println("Server listening!")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
