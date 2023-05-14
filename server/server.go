package server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nthr/files"
)

type Server struct {
	port      string
	mux       *http.ServeMux
	endpoints []string
	handlers  []func(w http.ResponseWriter, r *http.Request)
}

func Start() {
	port := ":3000"
	mux := http.NewServeMux()

	mux.HandleFunc("/sync", checkSync)
	mux.HandleFunc("/upload", downloadFile)

	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal("Server failed")
	}
}

func checkSync(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		sm := files.SyncMessage{}
		err := decoder.Decode(&sm)

		if err != nil {
			log.Fatal("could not decode request")
			return
		}
		fs := files.GetFileStructure(sm.Path)

		if sm.Hash == fs.Hash() {
			_, err := fmt.Fprintf(w, "File structure matches")
			if err != nil {
				return
			}
		} else {
			_, err := fmt.Fprintf(w, "file structure hashes conflict")
			if err != nil {
				return
			}
		}
	}
}

func downloadFile(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	fmt.Println(string(b))
}
