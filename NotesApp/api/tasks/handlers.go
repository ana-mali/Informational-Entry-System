package tasks

import (
	"encoding/json"
	"net/http"
	"strconv"

	"NotesApp/services"
)

//GET a list of all tasks
func listTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := services.ListTasks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

//POST a new task 
func addTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Name     string     `json:"name"`
		Priority *string    `json:"priority"`
		Due      *time.Time `json:"due"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if strings.TrimSpace(req.Name) == "" {
		http.Error(w, "name is required", http.StatusBadRequest)
		return
	}

	task, err := services.AddTask(
		req.Name,
		req.Priority,
		req.Due,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(task)
}


//DELETE a task usig an ID
func delete(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	if err := services.DeleteTask(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}