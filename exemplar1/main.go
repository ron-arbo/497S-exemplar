package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", MathServer)
	http.ListenAndServe(":8080", nil)
}

//MathServer is the function called once a request to the server is made. It will parse the URL and call the appropriate math function
func MathServer(w http.ResponseWriter, r *http.Request) {
	//Split the URL around "/" to get each argument
	args := strings.Split(r.URL.Path, "/")
	operation := args[1]

	//Check operation, call appropriate math function
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

//addNums will parse the URL and add each argument following the operation declaration by one another
func addNums(w http.ResponseWriter, r *http.Request) {
	//Split the URL around "/" to find arguments, slice to find operands only
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]

	//Initial value is first operand, convert to int
	sum, _ := strconv.Atoi(operands[0])

	//Loop through operands, applying operation to each
	for i := 1; i < len(operands); i++ {
		num, _ := strconv.Atoi(operands[i])
		sum += num
	}

	//Print to browser
	fmt.Fprintf(w, "Sum: %d", sum)
}

//addNums will parse the URL and multiply each argument following the operation declaration by one another
func multiplyNums(w http.ResponseWriter, r *http.Request) {
	//Split the URL around "/" to find arguments, slice to find operands only
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]

	//Initial value is first operand, convert to int
	product, _ := strconv.Atoi(operands[0])

	//Loop through operands, applying operation to each
	for i := 1; i < len(operands); i++ {
		num, _ := strconv.Atoi(operands[i])
		product *= num
	}

	//Print to browser
	fmt.Fprintf(w, "Product: %d", product)
}

//subtractNums will parse the URL and subtract each argument following the operation declaration by one another
func subtractNums(w http.ResponseWriter, r *http.Request) {
	//Split the URL around "/" to find arguments, slice to find operands only
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]

	//Initial value is first operand, convert to int
	difference, _ := strconv.Atoi(operands[0])

	//Loop through operands, applying operation to each
	for i := 1; i < len(operands); i++ {
		num, _ := strconv.Atoi(operands[i])
		difference -= num
	}

	//Print to browser
	fmt.Fprintf(w, "Difference: %d", difference)
}

//divideNums will parse the URL and divide each argument following the operation declaration by one another
func divideNums(w http.ResponseWriter, r *http.Request) {
	//Split the URL around "/" to find arguments, slice to find operands only
	args := strings.Split(r.URL.Path, "/")
	operands := args[2:]

	//Initial value is first operand, convert to int
	quotient, _ := strconv.Atoi(operands[0])

	//Loop through operands, applying operation to each
	for i := 1; i < len(operands); i++ {
		num, _ := strconv.Atoi(operands[i])
		quotient /= num
	}

	//Print to browser
	fmt.Fprintf(w, "Quotient: %d", quotient)
}
