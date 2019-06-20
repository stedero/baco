package srv

import (
	"basta/ravo/baco/model"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"
)

type serverWriter struct {
	w http.ResponseWriter
}

const maxFormParseMemorySizeBytes = 10 * 1024 * 1024

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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	form := `<html>
				<head>
					<meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
					<title>Baco - Bart's XML to JSON converter for Revo</title>
				</head>
				<body>
					<h2>Baco - Bart's XML to JSON converter for Revo</h2>
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
		defer reader.Close()
		err = model.Convert(reader, &serverWriter{w})
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to transform XML to JSON: %v", err)))
	}
}

func getReader(r *http.Request) (io.ReadCloser, error) {
	contentType := r.Header["Content-Type"]
	if contentType != nil && strings.HasPrefix(contentType[0], "multipart/form-data") {
		r.ParseMultipartForm(maxFormParseMemorySizeBytes)
		file := r.MultipartForm.File["upload"]
		if file == nil {
			return nil, errors.New("no content")
		}
		fileHeader := file[0]
		return fileHeader.Open()
	}
	return r.Body, nil
}

func (sw *serverWriter) Write(data []byte) (n int, err error) {
	sw.w.Header().Set("Content-Type", "application/json; charset=utf-8")
	sw.w.WriteHeader(http.StatusOK)
	return sw.w.Write(data)
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
