package app

import (
	"log"
	"file-upload-go/domain"
	"file-upload-go/handlers"
	"file-upload-go/services"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func dummyFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "It work. This is a dummy function.")
}

func StartApp() {
	// 1. Create the router
	router := mux.NewRouter()

	handler := handlers.NewUploadHandler(
		services.NewLocalUploadService(
			domain.NewLocalUpload(),
		))

	// 2. Add functions to router
	router.HandleFunc("/upload", handler.FileUpload).Methods(http.MethodPost)
	router.HandleFunc("/dummy", dummyFunc)
	router.HandleFunc("/", Hello)
	http.Handle("/", router)

	// 3. Start the server
	fmt.Println("Server started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}