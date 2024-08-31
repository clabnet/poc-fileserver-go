# Fileserver with GoLang

This application create a minimalist **fileserver** to serve static files, and also permit the **upload** files onto the `tempFiles` folder.

### Setup environment

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

### Test upload

Start a [simple web server](https://www.npmjs.com/package/http-server) on port 3000, then navigate to:

```
http://localhost:3000/frontend/fileupload.html
```

### Test fileserver

Open a browser to `http://localhost:8080/` to see the `tempFiles` content list

---

### References

- [Write a Lightweight API and Static File Server in Go](https://medium.com/swlh/write-a-lightweight-api-and-static-file-server-in-go-5e5b208ccdaf)
- [Building a Multipart-Form Upload Service in Go](https://medium.com/@mohitdubey_83162/building-a-multipart-form-upload-service-in-go-a-practical-guide-4f69069bc912)

---

Claudio Barca 31/8/2024
