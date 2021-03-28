package handlers

import (
	"encoding/json"
	"github.com/Mishanki/specialist-dz-1/internal/core"
	"github.com/Mishanki/specialist-dz-1/internal/models"
	"github.com/Mishanki/specialist-dz-1/pkg/helpers"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var GetItems = func(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	jsonString := helpers.MyReadFile()
	json.Unmarshal(jsonString, &items)
	w.Header().Set("Content-Type", "application/json")

	if len(items) == 0 {
		err := models.ErrorMsg{Success: false, Message: "No one items found in store back", Code: core.ITEMS_IS_NOT_FOUND_ERROR}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(items)
}

var GetItem = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var items []models.Item
	jsonString := helpers.MyReadFile()
	json.Unmarshal(jsonString, &items)
	w.Header().Set("Content-Type", "application/json")
	for _, v := range items {
		if v.Id == id {
			var item models.Item
			item = v
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	err := models.ErrorMsg{Success: false, Message: "Item with that id not found", Code: core.ITEM_IS_NOT_FOUND_ERROR}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)
}

var CreateItem = func(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	item.Id = id

	var items []models.Item
	jsonString := helpers.MyReadFile()
	json.Unmarshal(jsonString, &items)
	for _, v := range items {
		if item.Id == v.Id {
			err := models.ErrorMsg{Success: false, Message: "Item with that id already exists", Code: core.CREATE_ITEM_IS_EXIST}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(err)
			return
		}
	}

	items = append(items, item)
	file, _ := json.MarshalIndent(items, "", " ")
	helpers.MyWrite(file)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(models.SuccessMsg{Success: true, Message: "Item created"})
}

var UpdateItem = func(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	json.NewDecoder(r.Body).Decode(&item)

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var items []models.Item
	jsonString := helpers.MyReadFile()
	json.Unmarshal(jsonString, &items)
	for k, v := range items {
		if v.Id == id {
			items[k].Title = item.Title
			items[k].Amount = item.Amount
			items[k].Price = item.Price
			file, _ := json.MarshalIndent(items, "", " ")
			helpers.MyWrite(file)

			msg := models.SuccessMsg{Success: true, Message: "Item updated"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(msg)

			return
		}
	}

	err := models.ErrorMsg{Success: false, Message: "Item with that id not found", Code: core.UPDATE_ITEM_IS_NOT_FOUND}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(err)
}

var DeleteItem = func(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	var items []models.Item
	var newItems []models.Item
	jsonString := helpers.MyReadFile()
	json.Unmarshal(jsonString, &items)
	var isDeleted bool
	for _, v := range items {
		if v.Id != id {
			newItems = append(newItems, v)
		} else {
			isDeleted = true
		}
	}

	if isDeleted {
		file, _ := json.MarshalIndent(newItems, "", " ")
		helpers.MyWrite(file)

		msg := models.SuccessMsg{Success: true, Message: "Item was deleted"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(msg)
	} else {
		err := models.ErrorMsg{Success: false, Message: "Item with that id not found", Code: core.DELETE_ITEM_IS_NOT_FOUND}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}
}
