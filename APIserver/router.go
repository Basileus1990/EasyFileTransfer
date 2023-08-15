package main

import (
	"net/http"

	"github.com/Basileus1990/EasyFileTransfer.git/APIserver/receiver"
)

// API options
const (
	SHARE = "/share"
)

// returns configured mutex
func setRouting() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(SHARE, receiver.ShareFiles)

	return mux
}
