package domain

import (
	"fileserver/common"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

type IUpload interface {
	SaveFile(multipart.File, *multipart.FileHeader) (string, error)
}

type LocalUpload struct {
}

func (lu LocalUpload) SaveFile(file multipart.File, handler *multipart.FileHeader) (string, error) {

	// Retrieve file from form-data
	defer file.Close()

	fmt.Printf("\nFile name: %+v\n", handler.Filename)
	fmt.Printf("File size: %+v\n", handler.Size)
	fmt.Printf("File header: %+v\n", handler.Header)

	// make sure the temp folder exists
	tempFolderPath := fmt.Sprintf("%s%s", common.RootPath, "\\tempFiles")

	tempFile, err := os.Create(filepath.Join(tempFolderPath, handler.Filename))
	if err != nil {
		errStr := fmt.Sprintf("Error in creating the file %s\n", err)
		fmt.Println(errStr)
		return errStr, err
	}

	defer tempFile.Close()

	// Write upload file bytes to your new file
	filebytes, err := io.ReadAll(file)
	if err != nil {
		errStr := fmt.Sprintf("Error in reading the file buffer %s\n", err)
		fmt.Println(errStr)
		return errStr, err
	}

	tempFile.Write(filebytes)

	fmt.Printf("Successfully uploaded %s\n", handler.Filename)
	return "Successfully uploaded " + handler.Filename, nil
}

func NewLocalUpload() LocalUpload {
	return LocalUpload{}
}
