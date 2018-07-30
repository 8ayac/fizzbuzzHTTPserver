package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func viewHandler(w http.ResponseWriter, r *http.Request) {
	n, err := strconv.Atoi(r.URL.Path[10:])
	if err != nil {
		invalidNumber(w, r)
		return
	}

	fmt.Fprintf(w, "%q", fizzbuzz(n))
}

func fizzbuzz(n int) (result string) {
	if n%3 == 0 {
		result += "Fizz"
	}
	if n%5 == 0 {
		result += "Buzz"
	}
	if len(result) == 0 {
		result += strconv.Itoa(n)
	}
	return
}

func invalidNumber(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("invalid number"))
}

func main() {
	http.HandleFunc("/fizzbuzz/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
