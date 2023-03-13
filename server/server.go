package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/share", shareFiles)
	fmt.Println("<---- Starting the server ---->")
	http.ListenAndServe(":8080", nil)
}

func shareFiles(w http.ResponseWriter, r *http.Request) {
	log.Println("got to shareFiles")
	body, _ := ioutil.ReadAll(r.Body)
	log.Println(string(body))
}
