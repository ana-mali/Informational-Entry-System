package lists

import (
	"encoding/json"
	"net/http"
	"strconv"

	"NotesApp/services"
)

func Router() http.Handler {
	mux := http.NewServeMux()
	//Lists
	mux.HandleFunc("GET /lists", listLists)
	mux.HandleFunc("POST /lists/{name}", addList)
	mux.HandleFunc("DELETE /lists/{id}", deleteList)
	//Items 
	mux.HandleFunc("POST /lists/{listID}/items", addItem)
	mux.HandleFunc("DELETE /lists/{listID}/items/{itemID}", deleteItem)

	return mux
}
