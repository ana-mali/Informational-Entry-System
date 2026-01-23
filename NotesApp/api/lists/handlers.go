package lists

import (
	"encoding/json"
	"net/http"
	"strconv"

	"NotesApp/services"
)

// GET lists
func listLists(w http.ResponseWriter, r *http.Request) {
	lists, err := services.GetLists()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(lists)
}

//POST a new list
func addList(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name string `json:"name"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	list, err := services.CreateList(req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}

//DELETE a list using an ID
func deleteList(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := services.DeleteList(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// POST a new item to a list using a list ID and text for the item
func addItem(w http.ResponseWriter, r *http.Request) {
	listIDStr := r.PathValue("listID")
	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		http.Error(w, "invalid list id", http.StatusBadRequest)
		return
	}

	var req struct {
		Text string `json:"text"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	list, err := services.AddItem(listID, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(list)
}



//DELETE an item from a list using list ID and item ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	listIDStr := r.PathValue("listID")
	itemIDStr := r.PathValue("itemID")

	listID, err := strconv.Atoi(listIDStr)
	if err != nil {
		http.Error(w, "invalid list id", http.StatusBadRequest)
		return
	}

	itemID, err := strconv.Atoi(itemIDStr)
	if err != nil {
		http.Error(w, "invalid item id", http.StatusBadRequest)
		return
	}

	if err := services.RemoveItem(listID, itemID); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
