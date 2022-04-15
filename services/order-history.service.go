package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	m "github.com/gusgus-project/models"
	rep "github.com/gusgus-project/repository/order-history.repository"
)

func InsertOrderHistory(w http.ResponseWriter, r *http.Request) {
	fmt.Println("InsertOrderHistory")

	model := m.OrderHistory{}

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

// func GetOrderHistoryToday(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("GetOrder")

// 	res, err := rep.GetToday()
// 	if err != nil {
// 		WriteResponse(w, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	WriteResponse(w, http.StatusOK, res)
// }
