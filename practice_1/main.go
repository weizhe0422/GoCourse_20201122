package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})

	httpSvr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	httpSvr.ListenAndServe()
}
