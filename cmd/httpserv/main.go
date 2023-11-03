package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// Port we listen on.
const portNum string = ":8080"

func main() {
	logger := log.New(os.Stdout, "", log.Ltime+log.Ldate)
	logger.Println("Starting our simple http server.")

	// Registering our handler functions, and creating paths.
	http.Handle("/", loggerMiddleware(logger, http.HandlerFunc(Home)))
	http.Handle("/info", loggerMiddleware(logger, http.HandlerFunc(Info)))

	logger.Println("Started on port", portNum)

	// Spinning up the server.
	err := http.ListenAndServe(portNum, nil)
	if err != nil {
		logger.Fatal(err)
	}
}

func loggerMiddleware(logger *log.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Printf("[%s]%s ", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// Handler functions.
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Homepage")
}

func Info(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Info page")
}
