package app

import (
	"fileserver/domain"
	"fileserver/handlers"
	"fileserver/services"
	"fmt"
	"log"
	"net/http"
	"os"
)

func StartApp() {
	port := 8090

	// ============================================================================================
	// This is the part related to static file server.
	// https://medium.com/the-bug-shots/create-a-simple-fileserver-in-golang-9cd54453d373
	directoryPath := "../tempFiles"

	// Create the directory if not exists
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		err := os.Mkdir(directoryPath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}

	// Create a file server handler to serve the directory's contents
	fileServer := http.FileServer(http.Dir(directoryPath))

	// Create a new HTTP server and handle requests with CORS
	http.Handle("/", corsMiddleware(fileServer))

	// ============================================================================================
	// This is the part related to upload files.
	handler := handlers.NewUploadHandler(
		services.NewLocalUploadService(
			domain.NewLocalUpload(),
		))

	http.HandleFunc("/health", dummyFunc)

	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		handler.FileUpload(w, r)
	})

	// ============================================================================================
	// Start the server
	fmt.Printf("\nServer started at http://localhost:%d\n", port)
	fmt.Println("Try the upload at http://localhost:3000/frontend/fileupload.html")
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

func dummyFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to minimalist FileServer Go.")
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
}

// CORS middleware to apply CORS headers to a handler
func corsMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)
		h.ServeHTTP(w, r)
	})
}
