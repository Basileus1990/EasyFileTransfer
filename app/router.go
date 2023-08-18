package app

import (
	"github.com/Basileus1990/EasyFileTransfer.git/src/application/controllers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/files", controllers.AddDirInfo).Methods("POST")
	router.HandleFunc("/files", controllers.AddDirInfo).Methods("GET")

	return router
}
