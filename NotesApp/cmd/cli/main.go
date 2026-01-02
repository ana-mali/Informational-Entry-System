package main

import (
	"fmt"
	"os"
	"strconv"
	"NotesApp/services"
)

func main() {
	//Global Validation 
	if len(os.Args) < 2 ||len(os.Args) < 3{
		printHelp("No command provided.")
		return
	}
	app := os.Args[1]
	if app!="app"{
		printHelp(fmt.Sprintf("Unknown application: %s", app))
		return
	}

	//Command validations and executions
	cmd := os.Args[2]
	switch cmd {
	case "help", "--help":
		printHelp("")
		return

	case "add":
		if len(os.Args) < 4 {
			fmt.Println("Please provide note text.")
			return
		}
		
		note, err := services.AddNote(os.Args[3])
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Note added with ID:", note.ID)
		}

		

	case "list":
		notes, err := services.ListNotes()
		if err != nil {
			fmt.Println("Error:", err)
		}else if len(notes)==0 {
			fmt.Println("No notes found.")
		}else{
			for _, note := range notes {
				fmt.Printf("[%d] %s\n", note.ID, note.Text)
			}
		}

	case "delete":
		if len(os.Args) < 4  {
			fmt.Println("Please provide a note ID.")
			return
		}

		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("Note ID must be a number.")
			return
		}
		err = services.DeleteNote(id)
		if err != nil {
			fmt.Println("Error:", err)
		}else{
			fmt.Println("Note deleted with ID:", id)
		}
	default:
		printHelp(fmt.Sprintf("Unknown command: %s", cmd))
	}
}

func printHelp(msg string) {
	if msg != "" {
		fmt.Println(msg)
		fmt.Println()
	}

	fmt.Println(`Notes CLI — Personal Notes Manager

Usage:
  app <command> [arguments]

Commands:
  add <text>        Add a new note
  list              List all notes
  delete <id>       Delete a note
  help              Show this help

Examples:
  app add "Buy groceries"
  app list
  app delete 2
`)
}