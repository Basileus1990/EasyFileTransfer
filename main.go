package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/Basileus1990/EasyFileTransfer.git/app"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(0)
	}

	if strings.ToLower(os.Args[1]) == "server" {
		router := app.GetRouter()

		http.ListenAndServe(":8080", router)
	} else if strings.ToLower(os.Args[1]) == "files" {
		print("files")
	}
}
