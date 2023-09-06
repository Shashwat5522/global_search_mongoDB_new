package router

import (
	"golang_globalsearch_new/apis/search"
	"golang_globalsearch_new/apis/search/database"
	"golang_globalsearch_new/initializers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	Mgr      *initializers.Manager
	searchDB database.SearchDB
	handler  search.Handler
)

func init() {
	Mgr = initializers.GetInstance()
	searchDB = database.NewSearchDB(Mgr)
	handler = search.NewHandler(searchDB)

}
func Run() {
	router := mux.NewRouter()
	router.HandleFunc("/create", handler.CreateObjects).Methods("POST")
	router.HandleFunc("/search",handler.SearchObjects).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
