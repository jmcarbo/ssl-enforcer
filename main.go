package main

import (
	"fmt"
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("Redirecting to %s\n", "https://"+r.Host+r.URL.Path)
	http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusMovedPermanently)
}

func ok(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I am healthy!")
}

func main() {
	http.HandleFunc("/gce-health-check", ok)
	http.HandleFunc("/", redirect)
	log.Println("Listening on :9090")
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
