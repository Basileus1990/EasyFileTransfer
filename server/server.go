package server

import (
	"fmt"
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
