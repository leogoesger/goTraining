package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Profile is ..
type Profile struct {
	Name    string
	Hobbies []string
}

func cat(res http.ResponseWriter, req *http.Request) {
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
	http.HandleFunc("/leo", cat)
	http.ListenAndServe(":8080", nil)
}
