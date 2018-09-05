/*
ListenAndServe on port ":8080" using the default ServeMux.

Use HandleFunc to add the following routes to the default ServeMux:

"/"
"/dog/"
"/me/

Add a func for each of the routes.

Have the "/me/" route print out your name.
*/

package main

import (
	"fmt"
	"net/http"
)

func dog(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "dog")
}

func me(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Leo")
}

func others(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "others")
}

func main() {
	http.HandleFunc("/", others)
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.HandleFunc("/me/", me)
	http.ListenAndServe(":8080", nil)
}
