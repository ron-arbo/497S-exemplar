package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", HelloServer)
	// http.HandleFunc("/add", addNums)
	// http.HandleFunc("/subtract", subtractNums)
	// http.HandleFunc("/multiply", multiplyNums)
	// http.HandleFunc("/divide", divideNums)
	http.ListenAndServe(":8080", nil)
}

//MathServer greets the user
func MathServer(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[1:])
	args := strings.Split(r.URL.Path, "/")
	operation := args[1]
	fmt.Println(operation)
	fmt.Printf("%T\n", operation)
	switch operation {
	case "add":
		addNums(w, r)
	case "multiply":
		multiplyNums(w, r)
	case "subtract":
		subtractNums(w, r)
	case "divide":
		divideNums(w, r)
	}

}

//addNums
func addNums(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]
	fmt.Println("Operands: ", operands)
	sum := 0
	for _, s := range operands {
		num, _ := strconv.Atoi(s)
		sum += num
	}
	fmt.Println("Sum:", sum)
}

//multiplyNums
func multiplyNums(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]
	fmt.Println("Operands: ", operands)
	product := 0
	for _, s := range operands {
		num, _ := strconv.Atoi(s)
		product += num
	}
	fmt.Println("Product:", product)
}

//subtractNums
func subtractNums(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]
	fmt.Println("Operands: ", operands)
	difference := 0
	for _, s := range operands {
		num, _ := strconv.Atoi(s)
		difference += num
	}
	fmt.Println("Difference:", difference)
}

//divideNums
func divideNums(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]
	fmt.Println("Operands: ", operands)
	quotient := 0
	for _, s := range operands {
		num, _ := strconv.Atoi(s)
		quotient += num
	}
	fmt.Println("Quotient:", quotient)
}
