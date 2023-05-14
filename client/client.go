package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nthr/files"
)

func main() {
	path := "tempDir"
	fs := files.GetFileStructure(path)
	//sm := files.SyncMessage{
	//	Hash: fs.Hash(),
	//	Path: path,
	//}
    dm := files.DiffMessage{
        Path: path,
        Structure: fs.Stringify(),
    }
	body, err := json.Marshal(dm)
	if err != nil {
		log.Fatal("Could not marshal message")
	}

	resp, err := http.Post("http://127.0.0.1:3000/diff", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("L + bozo " + err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(b))
}
