package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Info struct {
	Start    time.Time
	HitCount int
	Version  string
}

var info = &Info{
	Start:    time.Now(),
	HitCount: 0,
	Version:  "version 1.0",
}
var mu sync.Mutex

func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return result
	}
	return 1
}

func FactorialHandler(w http.ResponseWriter, r *http.Request) {

	narg := mux.Vars(r)["narg"]

	log.Printf("Got %s\n", narg)

	n, err := strconv.Atoi(narg)
	if err != nil {
		log.Println(err)
		http.Error(w, fmt.Sprintf("Error n argument: %s", err.Error()), http.StatusBadRequest)
		return
	}
	if n < 0 {
		http.Error(w, fmt.Sprintf("Error n argument: %d not valid", n), http.StatusBadRequest)
		return
	}

	time.Sleep(1000 * time.Millisecond)

	result := Factorial(n)

	// LETS MAKE SURE TO INCREMENT HIT COUTNER ONLY AFTER RUNNING FACTORIAL
	// IN CASE WE FAIL ANYWHERE BEFORE
	// THIS WAY WE ONLY COUNT SUCCESSFULL CALLS
	mu.Lock()
	info.HitCount++
	mu.Unlock()

	json.NewEncoder(w).Encode(result)
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	elapsed := time.Since(info.Start)

	b := struct {
		UpSince  time.Time `json:"upsince"`
		UpTime   float64   `json:"uptime"`
		HitCount int       `json:"hitcount"`
		Version  string    `json:"version"`
	}{
		UpSince:  info.Start,
		UpTime:   elapsed.Seconds(),
		HitCount: info.HitCount,
		Version:  info.Version,
	}
	json.NewEncoder(w).Encode(b)
}

func main() {

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET"},
	})
	r := mux.NewRouter()

	r.HandleFunc("/info", InfoHandler).Methods("GET")
	r.HandleFunc("/factorial/{narg}", FactorialHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8081", c.Handler(r)))

}
