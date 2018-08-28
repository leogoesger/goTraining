package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type hotdog int

// Profile is ..
type Profile struct {
	Name    string
	Hobbies []string
}

func (h hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalf("Nop")
	}
	profile := Profile{"Alex", []string{"snowboarding", "programming"}}
	js, err := json.Marshal(profile)
	if err != nil {
		log.Fatalf("json failed")
	}

	fmt.Println(string(js))
	res.Header().Set("Content-Type", "application/json")
	res.Write(js)
}

func main() {
	var h hotdog
	http.ListenAndServe(":8080", h)
}
