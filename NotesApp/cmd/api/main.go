package main

import (
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "NotesApp/services"
)

func main() {
    http.HandleFunc("/notes", notesHandler)
    log.Println("API running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func notesHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        // POST for Insert New Note
        text := r.URL.Query().Get("text")
        note, err := services.AddNote(text)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(note)

    case http.MethodGet:
        // GET for list
        notes, err := services.ListNotes()
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        json.NewEncoder(w).Encode(notes)

    case http.MethodDelete:
        // DELETE for deletion of a note
        idStr := r.URL.Query().Get("id")
        id, _ := strconv.Atoi(idStr)
        err := services.DeleteNote(id)
        if err != nil {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        w.WriteHeader(http.StatusNoContent)

    default:
        http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
    }
}
