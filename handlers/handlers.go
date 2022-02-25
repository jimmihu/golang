package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleReq() {
	log.Println("Start development server localhost:9999")
	//bikin router dengan gorilla/mux
	myRouter := mux.NewRouter().StrictSlash(true)
	//routing
	myRouter.HandleFunc("/", HomePage)
	myRouter.HandleFunc("/user", CreateUser).Methods("OPTIONS", "POST")
	myRouter.HandleFunc("/users/{page}/{take}", GetUsers).Methods("OPTIONS", "GET")
	myRouter.HandleFunc("/user/{id}", GetDetail).Methods("OPTIONS", "GET")
	//serving
	log.Fatal(http.ListenAndServe(":9999", myRouter))
}
