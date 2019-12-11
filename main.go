package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
)

// 1. ECHO
func EchoHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s", err.Error())))
		return
	}

	response := Echo(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

// 2. INVERT
func InvertHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s", err.Error())))
		return
	}

	response := Invert(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

// 3. FLATTEN
func FlattenHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s", err.Error())))
		return
	}

	response := Flatten(records)
	fmt.Fprint(w, "\n"+response+"\n")
}

// 4. SUM
func SumHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s", err.Error())))
		return
	}

	response, err := Sum(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror computing sum: %s", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response+"\n")
}

// 5. MULTIPLY
func MultiplyHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("file")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror %s", err.Error())))
		return
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror reading CSV: %s", err.Error())))
		return
	}

	response, err := Multiply(records)
	if err != nil {
		w.Write([]byte(fmt.Sprintf("\nerror computing multiply: %s", err.Error())))
		return
	}
	fmt.Fprint(w, "\n"+response+"\n")
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
