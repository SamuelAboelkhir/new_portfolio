package main

import (
	"fmt"
	"log"
	"net/http"
)

func server(cfg *apiConfig) {
	fmt.Printf("Server started on port %s\n", cfg.port)
	mux := http.NewServeMux()

	mux.Handle("/app/", http.StripPrefix("/app/", logsMiddleware(http.FileServer(http.Dir(cfg.filePathRoot)))))
	mux.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir(cfg.publicPath))))

	mux.HandleFunc("GET /{$}", cfg.handlerHome)
	mux.HandleFunc("GET /greeting", cfg.handleGreeting)

	if err := http.ListenAndServe(":"+cfg.port, mux); err != nil {
		log.Fatal(err)
	}
}

func logsMiddleware(next http.Handler) http.Handler {
	fmt.Println("Logs middleware attached")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
