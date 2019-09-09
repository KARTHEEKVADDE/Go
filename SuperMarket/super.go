package super

import (
	"encoding/json"
	//"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Item struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Price string `json:"price,omitempty"`
}

var items []Item

func GetItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Item{})
}

func GetItems(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func CreateItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var item Item
	_ = json.NewDecoder(req.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func UpdateItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	var item Item
	_ = json.NewDecoder(req.Body).Decode(&item)
	item.ID = params["id"]
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func DeleteItem(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range items {
		if item.ID == params["id"] {
			items = append(items[:index], items[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(items)
}
