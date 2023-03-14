package server

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func StartServer() {
	mux := setRouting()

	fmt.Println("<---- Starting the server ---->")
	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}

func shareFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("got to shareFiles")
	body, _ := io.ReadAll(r.Body)
	log.Println(string(body))
}
