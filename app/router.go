package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Item represents the data structure for our resource.
type Item struct {
    ID   string `json:"id,omitempty"`
    Name string `json:"name,omitempty"`
}
// Database (for simplicity, we use an in-memory slice).
var items []Item

// GetItems returns all items in the database.
func GetItems(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(items)
}

// GetItem returns a specific item by ID.
func GetItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for _, item := range items {
        if item.ID == params["id"] {
            json.NewEncoder(w).Encode(item)
            return
        }
    }
    json.NewEncoder(w).Encode(&Item{})
}

// CreateItem adds a new item to the database.
func CreateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var item Item
    _ = json.NewDecoder(r.Body).Decode(&item)
    items = append(items, item)
    json.NewEncoder(w).Encode(item)
}

// UpdateItem updates an existing item in the database.
func UpdateItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range items {
        if item.ID == params["id"] {
            items[index] = Item{ID: params["id"], Name: "UpdatedName"}
            json.NewEncoder(w).Encode(items[index])
            return
        }
    }
    json.NewEncoder(w).Encode(&Item{})
}

// DeleteItem removes an item from the database.
func DeleteItem(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    params := mux.Vars(r)
    for index, item := range items {
        if item.ID == params["id"] {
            items = append(items[:index], items[index+1:]...)
            break
        }
    }
    json.NewEncoder(w).Encode(items)
}
