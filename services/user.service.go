package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/gusgus-project/models"
	rep "github.com/gusgus-project/repository/user.repository"
)

func InsertUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertUser")

	model := m.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := rep.Insert(model)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetUser")

	params := mux.Vars(r)

	id := params["id"]

	res, err := rep.Get(id)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func GetJoin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetJoin")

	params := mux.Vars(r)

	id := params["id"]

	res, err := rep.GetJoin(id)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchUser")

	var filter interface{}
	name := r.URL.Query().Get("name")
	jsonData := `{"name": { "$regex": ".*` + name + `.*" }, "isActive": 1 }`

	err := json.Unmarshal([]byte(jsonData), &filter)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := rep.Search(filter)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("UpdateUser  " + id)

	model := m.User{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = json.Unmarshal(body, &model)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	err = rep.Update(id, model)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, "")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("DeleteUser  " + id)

	err := rep.Delete(id)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, "")
}
