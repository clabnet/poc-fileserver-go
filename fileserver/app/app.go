package app

import (
	"log"
	"fileserver/domain"
	"fileserver/handlers"
	"fileserver/services"
	"fmt"
	"net/http"
	"os"
)

func StartApp() {
	port := 8080

	// ============================================================================================
	// This is the part related to static file serve. It should be refactor to a separate file ...
	// https://medium.com/the-bug-shots/create-a-simple-fileserver-in-golang-9cd54453d373
	directoryPath := "../tempFiles"

	// Check if the directory exists
	_, err := os.Stat(directoryPath)
	if os.IsNotExist(err) {
		fmt.Printf("Directory '%s' not found.\n", directoryPath)
		return
	}

	// Create a file server handler to serve the directory's contents
	fileServer := http.FileServer(http.Dir(directoryPath))

	// Create a new HTTP server and handle requests
	http.Handle("/", fileServer)

	// ============================================================================================
	// This is the part related to upload files. It is already refactorized code.
	handler := handlers.NewUploadHandler(
		services.NewLocalUploadService(
			domain.NewLocalUpload(),
		))

	http.HandleFunc("/health", dummyFunc)
	http.HandleFunc("/upload", handler.FileUpload)

	// ============================================================================================
	// Start the server
	fmt.Printf("Server started at http://localhost:%d\n", port)
	fmt.Println("Test upload at http://localhost:3000/frontend/fileupload.html")
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	log.Fatal(err)
}

func dummyFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, welcome to minimalist FileServer Go.")
}

