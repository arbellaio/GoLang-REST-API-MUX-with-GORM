package main

import (
	"github.com/gorilla/mux"
	"goLangRestApiMuxGorm/DbContext_Migration"
	"goLangRestApiMuxGorm/Services"
	"log"
	"net/http"
)

//To install gorilla mux
//go get -u github.com/gorilla/mux

//To install gorm (gorilla ORM)
//go get -u gorm.io/gorm

//To install gorm mysql connector
//go get -u gorm.io/driver/mysql

// app code starts here

//initializeRouter to create apis
func initializeRouter()  {
	router := mux.NewRouter()
	router.HandleFunc("/api/users", Services.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{id}", Services.GetUser).Methods("GET")
	router.HandleFunc("/api/users", Services.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/{id}", Services.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{id}", Services.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":9000", router))
}



func main() {
	dbContext.InitialMigration()
	initializeRouter()
}
