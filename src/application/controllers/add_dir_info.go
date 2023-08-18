package controllers

import "net/http"

func AddDirInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AddDirInfo"))
}
