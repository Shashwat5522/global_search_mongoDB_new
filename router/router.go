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
	
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",http.FileServer(http.Dir("./static/"))))
	router.HandleFunc("/index",handler.ShowIndexPage).Methods("GET")
	router.HandleFunc("/object",handler.ShowObject).Methods("GET")
	router.HandleFunc("/object_page",handler.ShowObjectPage).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
