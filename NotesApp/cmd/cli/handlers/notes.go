package handlers

import (
	"fmt"
	"NotesApp/services"
	"strconv"
)

func HandleNotes(cmd string, args []string) {
	switch cmd {

	case "add":
		if len(args) < 1 {
			fmt.Println("Please provide note text.")
			return
		}
		note, err := services.AddNote(args[0])
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Note added with ID:", note.ID)

	case "list":
		notes, err := services.ListNotes()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		if len(notes) == 0 {
			fmt.Println("No notes found.")
			return
		}
		for _, n := range notes {
			fmt.Printf("[%d] %s\n", n.ID, n.Text)
		}

	case "delete":
		if len(args) < 1 {
			fmt.Println("Please provide note ID.")
			return
		}
		id, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("ID must be a number.")
			return
		}
		if err := services.DeleteNote(id); err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Note deleted:", id)

	default:
		fmt.Println("Unknown notes command:", cmd)
	}
}
