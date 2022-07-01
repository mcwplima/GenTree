package handler

import (
	"encoding/json"
	"gentree/errors"
	"gentree/storage"
	"gentree/util"

	model "gentree/models"

	"io/ioutil"

	"net/http"

	"github.com/gorilla/mux"
)

//Relation struct
type Relation struct{}

//Get Relation retrive a person
func (e *Relation) Get(w http.ResponseWriter, r *http.Request) {
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

	Relation := storage.Relation()
	resp, err := Relation.RelationGet(objectid)
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	util.SendJSON(w, 200, resp)
}

//Add Gradient add a environment
func (e *Relation) Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	request := &model.Relation{}
	err := json.Unmarshal(body, request)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	storage, ok := storage.FromContext(r.Context())
	if !ok {
		errors.SendError(w, errors.Database)
		return
	}

	if request.Child == nil {
		er := &errors.Error{Description: "Bad Request, empty child", Status: 400}
		errors.SendError(w, er)
		return
	}

	Person := storage.Person()
	resp, err := Person.PersonGet(*request.Child)
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	if request.Parent != nil {
		resp, err = Person.PersonGet(*request.Parent)
		if resp == nil && err == nil {
			errors.SendError(w, errors.NotFound)
			return
		}
		if err != nil {
			er := &errors.Error{Description: err.Error(), Status: 400}
			errors.SendError(w, er)
			return
		}
	}

	Relation := storage.Relation()
	objectid, err := Relation.RelationAdd(request)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	util.SendJSON(w, 200, objectid)
}

//Del Gradient delete a environment
func (e *Relation) Del(w http.ResponseWriter, r *http.Request) {
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

	Relation := storage.Relation()
	resp, err := Relation.RelationGet(objectid)
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}

	err = Relation.RelationDel(resp.ObjectID)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	errors.SendError(w, &errors.Error{Description: "OK", Status: 200})
}

//Update Gradient updates a gateway
func (e *Relation) Update(w http.ResponseWriter, r *http.Request) {
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

	Relation := storage.Relation()
	resp, err := Relation.RelationGet(objectid)
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	body, _ := ioutil.ReadAll(r.Body)
	request := &model.Relation{}
	err = json.Unmarshal(body, request)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}
	if request.Parent != nil {
		resp.Parent = request.Parent
	}
	if request.Child != nil {
		resp.Child = request.Child
	}

	_, err = Relation.RelationAdd(resp)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	errors.SendError(w, &errors.Error{Description: "OK", Status: 200})
}

//List Gradient retrive all gateways
func (e *Relation) List(w http.ResponseWriter, r *http.Request) {
	storage, ok := storage.FromContext(r.Context())
	if !ok {
		errors.SendError(w, errors.Database)
		return
	}

	Relation := storage.Relation()
	resp, err := Relation.RelationList()
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}
	util.SendJSON(w, 200, resp)
}
