package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"nthr/files"
	"os"
)

func main() {
	path := "tempDir"
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
	sendFile("tempdir/bleh")
}

func sendFile(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal("Could not load ", path)
	}
	msg := files.FileMessage{
		Key:  path,
		File: file,
	}
	toSend, err := json.Marshal(msg)
	res, err := http.Post("http://127.0.0.1:3000/upload", "application/json", bytes.NewBuffer(toSend))
	b, _ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		log.Fatal("Could not create request")
	}
	fmt.Printf("response code: ", string(b))
}
