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
	path := "testDir"
	fs := files.GetFileStructure(path)

	sm := files.SyncMessage{
		Hash: fs.Hash(),
		Path: path,
	}
	body, err := json.Marshal(sm)
	if err != nil {
		log.Fatal("Could not marshal message")
	}

	resp, err := http.Post("http://127.0.0.1:3000/sync", "application/json", bytes.NewBuffer(body))
	if err != nil {
		log.Fatal("L + bozo " + err.Error())
	}
	b, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(b))
}
