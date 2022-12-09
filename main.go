package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	app "github.com/parikalp/app"
)

// func main() {
// 	fmt.Println("Docker is Running")
// 	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 	// 	fmt.Fprintf(w, "Hello-World")
// 	// })
// 	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Hello, you've requested: %s with token: %s\n", r.URL.Path, r.URL.Query().Get("token"))
// 	})

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", rootHandler).Methods("GET")
	r.HandleFunc("/health", handlerFunc).Methods("GET")
	r.HandleFunc("/details", morehandle).Methods("GET")

	fmt.Println("Application Running on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))
}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	respone := map[string]string{
		"Name": "Git-hub",
		"timestamp":  time.Now().String(),
	}

	json.NewEncoder(w).Encode(respone)
}

func morehandle(w http.ResponseWriter, r *http.Request) {
	log.Println("Fetching the details")
	hostname, err := app.GetHostname()
	if err != nil {
		panic(err)
	}
	Ip, _ := app.GetIP()

	fmt.Println(hostname, Ip)
	respone := map[string]string{
		"hostname": hostname,
		"ip":       Ip.String(),
	}
	json.NewEncoder(w).Encode(respone)

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving the homepage")
	w.WriteHeader(http.StatusOK)
	fmt.Println(w, "Application is up and running")
}
