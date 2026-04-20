// Full Go Code (File Server)
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// =========================
// CONFIG
// =========================
const uploadDir = "./uploads"

// =========================
// INIT UPLOAD DIR
// =========================
func init() {
	os.MkdirAll(uploadDir, 0755)
}

// =========================
// UPLOAD HANDLER
// =========================
func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse multipart form (10 MB limit)
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "File not found", http.StatusBadRequest)
		return
	}
	defer file.Close()

	filePath := filepath.Join(uploadDir, handler.Filename)

	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Unable to create file", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully: %s\n", handler.Filename)
}

// =========================
// LIST FILES HANDLER
// =========================
func listFilesHandler(w http.ResponseWriter, r *http.Request) {
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		http.Error(w, "Unable to read directory", http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		fmt.Fprintln(w, f.Name())
	}
}

// =========================
// DOWNLOAD HANDLER
// =========================
func downloadHandler(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		http.Error(w, "File parameter missing", http.StatusBadRequest)
		return
	}

	filePath := filepath.Join(uploadDir, fileName)

	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/octet-stream")

	io.Copy(w, file)
}

// =========================
// MAIN FUNCTION
// =========================
func main() {

	http.HandleFunc("/upload", uploadHandler)
	http.HandleFunc("/files", listFilesHandler)
	http.HandleFunc("/download", downloadHandler)

	fmt.Println("Server started at http://localhost:8080")
	fmt.Println("Upload:   POST /upload (form-data key: file)")
	fmt.Println("List:     GET  /files")
	fmt.Println("Download: GET  /download?file=name")

	http.ListenAndServe(":8080", nil)
}