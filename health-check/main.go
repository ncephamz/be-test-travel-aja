package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func healthCheck(w http.ResponseWriter, r *http.Request){
   timeout := time.Duration(1 * time.Second)

	health := make(map[string]string)

	// update value here
	ip := "108.136.184.155:8000"

	_, err := net.DialTimeout("tcp", ip, timeout)
	if err != nil {
		fmt.Printf(err.Error())
		health["backend"] = "Error"
	} else {
		health["backend"] = "Ok"
	}

	bytes, err := json.Marshal(health)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(bytes)
}

func handleRequests() {
    http.HandleFunc("/health-check", healthCheck)
    log.Fatal(http.ListenAndServe(":8000", nil))
}

func main() {
    handleRequests()
}