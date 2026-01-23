package main

import (
	"log"
	"net/http"

	"NotesApp/api/notes"
	"NotesApp/api/tasks"
	"NotesApp/api/lists"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NotesApp API is running"))
	})

	mux.Handle("/notes/", notes.Router())
	mux.Handle("/tasks/", tasks.Router())
	mux.Handle("/lists/", lists.Router())

	log.Println("API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
