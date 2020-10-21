package service

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/felipesulzbach/exemplo-api-rest/app/src/model"
	"github.com/felipesulzbach/exemplo-api-rest/app/src/repository"
	"github.com/gorilla/mux"

)

// FindAllTeacher ...
func FindAllTeacher(w http.ResponseWriter, r *http.Request) {
	list, err := repository.FindAllTeacher()
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	for _, item := range list {
		log.Println(item.ToString())
	}

	jsonOkResponse(w, list)
}

// FindByIDTeacher ...
func FindByIDTeacher(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}

	entity, err := repository.FindByIDTeacher(id)
	switch {
	case err == sql.ErrNoRows:
		var errorDesc bytes.Buffer
		errorDesc.WriteString("ERROR: No records found for id=")
		errorDesc.WriteString(strconv.FormatInt(id, 10))
		log.Println(errorDesc.String())
		json.NewEncoder(w).Encode(errorDesc.String())
		return
	case err != nil:
		log.Panic(err)
		return
	default:
	}

	log.Println(entity.ToString())
	jsonOkResponse(w, entity)
}

// InsertTeacher ...
func InsertTeacher(w http.ResponseWriter, r *http.Request) {
	var entity model.Teacher
	_ = json.NewDecoder(r.Body).Decode(&entity)

	id, err := repository.InsertTeacher(entity)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
	jsonCreatedResponse(w, id)
}

// UpdateTeacher ...
func UpdateTeacher(w http.ResponseWriter, r *http.Request) {
	var entity model.Teacher
	_ = json.NewDecoder(r.Body).Decode(&entity)

	if err := repository.UpdateTeacher(entity); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Panic(err)
		return
	}
}

// DeleteTeacher ...
func DeleteTeacher(w http.ResponseWriter, r *http.Request) {
	// TODO
}
