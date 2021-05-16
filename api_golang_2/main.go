package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/getusers", returnAllUsers).Methods("GET")
	router.HandleFunc("/edituser", updateUsersMultipart).Methods("PUT")
	router.HandleFunc("/deleteuser", deleteUsersMultipart).Methods("DELETE")
	router.HandleFunc("/insertuser", insertUsersMultipart).Methods("POST")
	http.Handle("/", router)
	fmt.Println("Connected to port 1234")
	log.Fatal(http.ListenAndServe(":1234", router))

}
