package files

type SyncMessage struct {
	Hash string `json:"hash"`
	Path string `json:"path"`
}

type DiffMessage struct {
    Path string `json:"path"`
    Structure string `json:"structure"`
}
