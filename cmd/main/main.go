package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	//"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/Namozov/go-bookstore/pkg/routes"
)
func main(){
	r := mux.NewRouter()
	routes.RegisterBookStoreRouters(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}