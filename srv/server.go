package srv

import (
	"basta/ravo/baco/model"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
)

const maxFormParseMemorySizeBytes = 1024 * 1024

// Start the server
func Start(port int) {
	server := http.Server{Addr: ":" + strconv.Itoa(port)}
	http.HandleFunc("/", handle)
	server.ListenAndServe()
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		showForm(w)
	case "POST":
		process(w, r)
	default:
	}
}

func showForm(w http.ResponseWriter) {
	form := `<html>
				<head>
					<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
					<title>Baco - Bart's XML to JSON converter for Revo</title>
				</head>
				<body>
					<h1>Baco - Bart's XML to JSON converter for Revo</h1>
					<form action="/" method="post" enctype="multipart/form-data">
						<input type="file" name="upload"><br><br>
						<input type="submit">
					</form>
				</body>
			</html>`
	w.Write([]byte(form))
}

func process(w http.ResponseWriter, r *http.Request) {
	defer serverError(w, r)
	reader, err := getReader(r)
	if err == nil {
		rec := model.ReadRavoRecord(reader)
		if err == nil {
			w.Header().Set("Content-Type", "application/json; charset=utf-8")
			w.WriteHeader(200)
			rec.WriteJSON(w)
		}
	}
	if err != nil {
		msg := fmt.Sprintf("failed to transform: %v", err)
		w.WriteHeader(500)
		w.Write([]byte(msg))
	}
}

func getReader(r *http.Request) (io.Reader, error) {
	contentType := r.Header["Content-Type"]
	if contentType != nil && strings.HasPrefix(contentType[0], "multipart/form-data") {
		r.ParseMultipartForm(maxFormParseMemorySizeBytes)
		fileHeader := r.MultipartForm.File["upload"][0]
		return fileHeader.Open()
	}
	return r.Body, nil
}

func serverError(w http.ResponseWriter, rec *http.Request) {
	if r := recover(); r != nil {
		msg := fmt.Sprintf("%v", r)
		debug.PrintStack()
		w.Header().Set("Content-Type", "text/plain; charset=UTF-8")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(msg))
	}
}
