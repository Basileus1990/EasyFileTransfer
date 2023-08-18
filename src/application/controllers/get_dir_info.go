package controllers

import "net/http"

func GetDirInfo(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetDirInfo"))
}
