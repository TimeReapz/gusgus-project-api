package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	m "github.com/gusgus-project/models"
	rep "github.com/gusgus-project/repository/product.repository"
)

func InsertProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertProduct")

	model := m.Product{}

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

	fmt.Println(model)

	res, err := rep.Insert(model)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func GetProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetProduct")

	params := mux.Vars(r)

	id := params["id"]

	res, err := rep.Get(id)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func SearchProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println("SearchProduct")

	var filter interface{}
	name := r.URL.Query().Get("name")

	jsonData := `{"name": { "$regex": ".*` + name + `.*" }, "isActive": 1 }`

	err := json.Unmarshal([]byte(jsonData), &filter)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(filter)

	res, err := rep.Search(filter)
	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, res)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("UpdateProduct  " + id)

	model := m.Product{}

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

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	fmt.Println("DeleteProduct  " + id)

	err := rep.Delete(id)

	if err != nil {
		WriteResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	WriteResponse(w, http.StatusOK, "")
}
