package routes

import(
	
	"github.com/gorilla/mux"
	"github.com/Namozov/go-bookstore/pkg/controllers" //"./controllers" shunaqa qilsaman bo'ladi
)

var RegisterBookStoreRouters = func (router *mux.Router)  {
	router.HandleFunc("/book/", controller.CreateBook).Methods("POSt")
	router.HandleFunc("/book/", controller.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controller.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controller.DeleteBook).Methods("DELETE")
}