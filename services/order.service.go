package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/gusgus-project/models"
	rep "github.com/gusgus-project/repository/order.repository"
)

func InsertOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertOrder")

	model := m.Order{}

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

func GetOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetOrder")

	params := mux.Vars(r)

	id := params["id"]

	res, err := rep.Get(id)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func SearchOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchOrder")

	var filter interface{}
	name := r.URL.Query().Get("name")
	jsonData := `{"user.name": { "$regex": ".*` + name + `.*" }, "isActive": 1 }`

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

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("UpdateOrder  " + id)

	model := m.Order{}

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

func DeleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("DeleteOrder  " + id)

	err := rep.Delete(id)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, "")
}
