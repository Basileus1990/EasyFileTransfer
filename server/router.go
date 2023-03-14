package server

import "net/http"

const (
	SHARE = "/share"
)

// returns configured mutex
func setRouting() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(SHARE, shareFiles)

	return mux
}
