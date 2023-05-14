package files

import "os"

type SyncMessage struct {
	Hash string `json:"hash"`
	Path string `json:"path"`
}

type FileMessage struct {
	Key  string   `json:"key"`
	File *os.File `json:"file"`
}
