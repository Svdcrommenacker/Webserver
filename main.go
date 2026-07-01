package main

import (
	"log"
	"net/http"
)

func main() {

	const port = "8080"
	const filepathRoot = "."

	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir(filepathRoot))
	mux.Handle("/app/", http.StripPrefix("/app", fileServer))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	log.Println("Serving files from", filepathRoot, "on port:", port)
	log.Fatal(srv.ListenAndServe())
}
