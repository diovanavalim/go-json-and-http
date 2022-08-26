package main

import (
	"fmt"
	"http/json"
	"log"
	"net/http"
)

func c1(w http.ResponseWriter, r *http.Request) {
	json.Marshal()
	w.Write([]byte("Marshalling data!"))
}

func c2(w http.ResponseWriter, r *http.Request) {
	json.Unmarshal()
	w.Write([]byte("Unmarshalling data!"))
}

func main() {
	http.HandleFunc("/marshal", c1)
	http.HandleFunc("/unmarshal", c2)

	fmt.Println("Server Started")
	log.Fatal(http.ListenAndServe(":6060", nil))
}
