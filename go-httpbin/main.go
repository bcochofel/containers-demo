package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("GET Request")
	resp, err := http.Get("https://httpbin.org/get")
	if err != nil {
		log.Printf("The HTTP request failed with error %s\n", err)
	} else {
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Reading data failed with error %s\n", err)
		} else {
			fmt.Fprintf(w, string(data))
		}
	}
}

func main() {
	log.Println("Running on http://127.0.0.1:8080/")
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
