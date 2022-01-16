package main

import (
	"errors"
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the home page")
}

func About(w http.ResponseWriter, r *http.Request) {
	sum := addValues(2, 2)
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 plus 2 is %d", sum))
}

func addValues(x, y int) int {
	return x + y
}

func Divide(w http.ResponseWriter, r *http.Request){
	result, err := DivideNumbers(100.0, 0.0)
	if err != nil {
		fmt.Fprint(w, "Cannot divide by 0")
		return
	}

	fmt.Fprintf(w, fmt.Sprintf("The result of this division is %f", result))
}

func DivideNumbers(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot divide by zero")
		return 0, err
	}
	
	result := x / y 
	return result, nil
}

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))	
	_ = http.ListenAndServe(portNumber, nil)
}