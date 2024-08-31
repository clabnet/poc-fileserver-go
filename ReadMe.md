# UPLOAD with GoLang

### Setup environment

```
go get -u github.com/gorilla/mux
```

##### Install LiveReload

https://github.com/bokwoon95/wgo

```
go install github.com/bokwoon95/wgo@latest
```

### Start development server

Run main.go whenever a .go file changes.

```
cd fileserver
wgo run main.go
```

### Upload test

Start a simple web server on port 3000, then navigate to:

```
http://localhost:3000/frontend/fileupload.html
```

### Fileserver test

Open a browser to `localhost:8080/` to see the `tempFiles` content list

### References

[Write a Lightweight API and Static File Server in Go](https://medium.com/swlh/write-a-lightweight-api-and-static-file-server-in-go-5e5b208ccdaf)

- https://medium.com/@mohitdubey_83162/building-a-multipart-form-upload-service-in-go-a-practical-guide-4f69069bc912
