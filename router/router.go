package router

import (
	"github.com/gorilla/mux"

	"gentree/handler"
)

func Routes() *mux.Router {
	_Person := handler.Person{}
	_Relation := handler.Relation{}
	_Tree := handler.Tree{}

	// Router setup
	router := mux.NewRouter()
	router.StrictSlash(true)

	router.HandleFunc("/api/person/", _Person.Add).Methods("POST")
	router.HandleFunc("/api/person/{id}", _Person.Del).Methods("DELETE")
	router.HandleFunc("/api/person/{id}", _Person.Get).Methods("GET")
	router.HandleFunc("/api/person/{id}", _Person.Update).Methods("PUT")
	router.HandleFunc("/api/person/", _Person.List).Methods("GET")

	router.HandleFunc("/api/relation/", _Relation.Add).Methods("POST")
	router.HandleFunc("/api/relation/{id}", _Relation.Del).Methods("DELETE")
	router.HandleFunc("/api/relation/{id}", _Relation.Get).Methods("GET")
	router.HandleFunc("/api/relation/{id}", _Relation.Update).Methods("PUT")
	router.HandleFunc("/api/relation/", _Relation.List).Methods("GET")

	router.HandleFunc("/api/tree/{id}", _Tree.Get).Methods("GET")

	return router
}
