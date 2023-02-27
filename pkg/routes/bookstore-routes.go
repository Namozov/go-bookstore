package routes

import(
	
	"github.com/gorilla/mux"
	"github.com/Namozov/go-bookstore/pkg/controllers" //"./controllers" shunaqa qilsaman bo'ladi
)

var RegisterBookStoreRouters = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POSt")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.Delete.Book).Methods("DELETE")
}