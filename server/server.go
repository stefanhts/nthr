package server

import (
	"encoding/json"
	"fmt"
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
    mux.HandleFunc("/sync", checkSync) // TODO: rename this to syncStatus or something
    mux.HandleFunc("/diff", fileDiff)
	http.ListenAndServe(port, mux)
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

func fileDiff(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        decoder := json.NewDecoder(r.Body)
        dm := files.DiffMessage{}
        err := decoder.Decode(&dm)
        if err != nil {
            log.Fatal("could not decode fileDiff request")
            return
        }
        fs := files.GetFileStructure(dm.Path)
        fs.Display()
        serverSideFiles := fs.Stringify()
        clientSideFiles := dm.Structure
        fmt.Printf("server side files:\n%s\n\n", serverSideFiles)
        fmt.Printf("client side files:\n%s\n", clientSideFiles)
    }
}
