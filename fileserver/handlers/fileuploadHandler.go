package handlers

import (
	"log"
	"fileserver/services"
	"fmt"
	"net/http"
	"encoding/json"
)

type UploadHandler struct {
	service services.IUploadService
}

func (uh UploadHandler) FileUpload(w http.ResponseWriter, r *http.Request) {

	// Limit the size of the uploaded file (10 MB)
	r.ParseMultipartForm(10 << 20)

	// Get the file from the request
	file, handler, err := r.FormFile("file")
	if err != nil {
		errStr := fmt.Sprintf("Error in FileUpload.\n %s\n", err)
		fmt.Println(errStr)
		fmt.Fprintf(w, errStr)
		return
	}
	defer file.Close()

	result, err := uh.service.SaveFile(file, handler)

	if err != nil {
		log.Fatal(err)
		return
	}

	// Set the response content type to application/json
	w.Header().Set("Content-Type", "application/json")

	// Encode the response as JSON and send it
	if err := json.NewEncoder(w).Encode(map[string]string{"message": result}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func NewUploadHandler(service services.IUploadService) UploadHandler {
	return UploadHandler{
		service: service,
	}
}
