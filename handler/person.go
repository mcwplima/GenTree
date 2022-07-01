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

//Person struct
type Person struct{}

//Get Person retrive a person
func (e *Person) Get(w http.ResponseWriter, r *http.Request) {
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

	Person := storage.Person()
	resp, err := Person.PersonGet(objectid)
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
func (e *Person) Add(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	request := &model.Person{}
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

	if request.Name == nil {
		er := &errors.Error{Description: "Bad Request, empty name", Status: 400}
		errors.SendError(w, er)
		return
	}

	Person := storage.Person()
	objectid, err := Person.PersonAdd(request)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	util.SendJSON(w, 200, objectid)
}

//Del Gradient delete a environment
func (e *Person) Del(w http.ResponseWriter, r *http.Request) {
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

	Person := storage.Person()
	resp, err := Person.PersonGet(objectid)
	if resp == nil && err == nil {
		errors.SendError(w, errors.NotFound)
		return
	}

	err = Person.PersonDel(resp.ObjectID)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	errors.SendError(w, &errors.Error{Description: "OK", Status: 200})
}

//Update Gradient updates a gateway
func (e *Person) Update(w http.ResponseWriter, r *http.Request) {
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

	Person := storage.Person()
	resp, err := Person.PersonGet(objectid)
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
	request := &model.Person{}
	err = json.Unmarshal(body, request)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}
	if request.Name != nil {
		resp.Name = request.Name
	}

	_, err = Person.PersonAdd(resp)
	if err != nil {
		er := &errors.Error{Description: err.Error(), Status: 400}
		errors.SendError(w, er)
		return
	}

	errors.SendError(w, &errors.Error{Description: "OK", Status: 200})
}

//List Gradient retrive all gateways
func (e *Person) List(w http.ResponseWriter, r *http.Request) {
	storage, ok := storage.FromContext(r.Context())
	if !ok {
		errors.SendError(w, errors.Database)
		return
	}

	Person := storage.Person()
	resp, err := Person.PersonList()
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
