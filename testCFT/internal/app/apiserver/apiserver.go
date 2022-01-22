package apiserver

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Apiserver struct {
	router *mux.Router
}
type FileWithData struct {
	name string
	data []byte
}

func New() *Apiserver {
	return &Apiserver{
		router: mux.NewRouter(),
	}
}

func (s *Apiserver) ConfigureApi() {
	s.router.HandleFunc("/getfile", getFile).Methods("GET")
	s.router.HandleFunc("/getfile/{name}", getFiles).Methods("GET")
	s.router.HandleFunc("/sendfile", createFile).Methods("PUT")
	s.router.HandleFunc("/sendfile", updateFile).Methods("POST")
	s.router.HandleFunc("/deletefile/{name}", deleteFile).Methods("DELETE")
}

func (s *Apiserver) Start() error {
	s.ConfigureApi()
	return nil
}

func getFile(w http.ResponseWriter, r *http.Request) {
	file, error := os.ReadFile(r.URL.Query().Get("name"))
	if error != nil {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(file)
	// base64str := base64.StdEncoding.EncodeToString(file)

	// w.Header().Set("Content-Type", "application/json")
	// file, err := os.ReadFile(r.bo)
}
func getFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	files, err := os.ReadDir("tmp")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	result := []FileWithData{}
	for _, file := range files {
		hash, err := GetHash(file.Name())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		temp := FileWithData{
			name: file.Name(),
			data: hash,
		}
		result = append(result, temp)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
func createFile(w http.ResponseWriter, r *http.Request) {
	// file := FileWithData{}
	// json.NewDecoder(r.Body).Decode(file)
	// os.WriteFile("/tmp/"+file.name, file.data, os.ModeAppend)
}
func updateFile(w http.ResponseWriter, r *http.Request) {

}
func deleteFile(w http.ResponseWriter, r *http.Request) {

}
func GetHash(path string) (hash []byte, err error) {
	f, ferr := os.Open(path)
	if ferr != nil {
		err = ferr
		return
	}
	defer f.Close()
	h := sha256.New()
	if _, herr := io.Copy(h, f); herr != nil {
		err = herr
		return
	}
	hash = h.Sum(nil)
	return
}
func GetHashBytes(bytes []byte) (hash []byte, err error) {
	h := sha256.New()
	hash = h.Sum(bytes)
	return
}
