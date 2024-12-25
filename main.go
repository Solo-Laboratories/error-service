package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	// Define the file server for static content
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	// Define a custom 404 handler
	// http.HandleFunc("/404.html", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "static/404.html")
	// })

	// Start the server
	port := "8080"
	slog.Info(fmt.Sprintf("Starting server on port %s\n", port))
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		slog.Error(fmt.Sprintf("Error while starting webserver: %v", err))
		panic("unable to start webserver")
	}
}
