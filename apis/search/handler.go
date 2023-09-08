package search

import (
	"encoding/json"
	"golang_globalsearch_new/apis/search/database"
	"html/template"
	"log"
	"net/http"
)

type Handler struct {
	searchDB database.SearchDB
}

func NewHandler(searchDB database.SearchDB) Handler {
	return Handler{
		searchDB: searchDB,
	}
}

func (h *Handler) CreateObjects(w http.ResponseWriter, r *http.Request) {

	err := h.searchDB.CreateObjects()
	if err != nil {
		log.Fatalf("error is:=%v\n", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("objects created successfully!!!"))
}

func (h *Handler) SearchObjects(w http.ResponseWriter, r *http.Request) {
	searchWord := r.URL.Query()
	resp, err := h.searchDB.SearchObject(searchWord.Get("search"))
	if err != nil {
		log.Fatal(err)
	}
	response, jsonErr := json.Marshal(resp)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)
	}

	func (h *Handler) ShowIndexPage(w http.ResponseWriter, r *http.Request) {
		tmpl, err := template.ParseFiles("C:/Users/Bacancy/Desktop/golang_practice/golang_globalsearch_new/templates/index.html")
		if err != nil {
			log.Fatal(err)
		}
		if err := tmpl.Execute(w, nil); err != nil {
			log.Fatal(err)
		}

	}
	func(h *Handler)ShowObject(w http.ResponseWriter,r *http.Request){
		objectID:=r.Header.Get("object-Id")
		log.Println(objectID)
		resp,err:=h.searchDB.ShowObject(objectID)
		if err!=nil{
			log.Fatalf("error is %v\n",err)
		}
		response,jsonErr:=json.Marshal(resp)
		if jsonErr!=nil{
			log.Fatal(jsonErr)
		}
		w.Header().Set("Content-Type","application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(response)

	}

	func(h *Handler)ShowObjectPage(w http.ResponseWriter,r *http.Request){
		tmpl,err:=template.ParseFiles("C:/Users/Bacancy/Desktop/golang_practice/golang_globalsearch_new/templates/object.html")
		if err!=nil{
			log.Fatal(err)
		}
		if err:=tmpl.Execute(w,nil);err!=nil{
			log.Fatal(err)
		}
	}

