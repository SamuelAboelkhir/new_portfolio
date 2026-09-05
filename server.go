package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func server() {
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	fmt.Printf("Test server started on port %s\n", port)
	http.ListenAndServe(":"+port, testMiddleware(http.Dir("./app")))
}

func testMiddleware(root http.FileSystem) http.Handler {
	fmt.Println("Middleware attached")
	handler := http.FileServer(root)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("%s %s", r.Method, r.URL.Path)
		handler.ServeHTTP(w, r)
	})
}
