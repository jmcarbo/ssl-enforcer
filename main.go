package main

import (
	"log"
	"net/http"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	log.Printf("hello from redirect %s\n", "https://"+r.Host+r.URL.Path)
	http.Redirect(w, r, "https://"+r.Host+r.URL.Path, http.StatusMovedPermanently)
}

func main() {
	http.HandleFunc("/", redirect)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}
