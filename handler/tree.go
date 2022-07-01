package handler

import (
	"gentree/errors"
	"gentree/storage"
	"gentree/util"

	"net/http"

	"github.com/gorilla/mux"
)

//Tree struct
type Tree struct{}

//Get Relation retrive a person
func (e *Tree) Get(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	vars := mux.Vars(r)
	objectid, ok := vars["id"]
	if !ok {
		errors.SendError(w, errors.BadRequest)
		return
	}
	storage, ok := storage.FromContext(r.Context())
	if !ok {
		errors.SendError(w, errors.Database)
		return
	}

	Tree := storage.Tree()
	tree, err := Tree.TreeGet(objectid)
	if tree == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	Person := storage.Person()

	for index, p := range tree {
		relations, err := Person.PersonListByChild(p.ObjectID)
		if err != nil {
			er := &errors.Error{Description: err.Error(), Status: 400}
			errors.SendError(w, er)
			return
		}

		tree[index].Relations = relations

	}
	util.SendJSON(w, 200, tree)
}
