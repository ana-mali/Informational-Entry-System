package main

import (
	"fmt"
	"os"
)
//const cmds := []string{"add", "list", "view", "delete"}

func main() {
	//Global Validation and help flags
	if os.Args[1]=="help" || os.Args[1]=="==help" {
		fmt.Println(`Notes CLI — Personal Notes Manager

Usage:
  notes <command> [arguments]

Available Commands:
  add <text>        Add a new note
  list              List all notes
  view <id>         View a note by its ID
  delete <id>       Delete a note by its ID
  help              Show this help message

Examples:
  notes add "Buy groceries"
  notes list
  notes view 1
  notes delete 2

Notes:
  - Each command runs once and then exits.
  - Note IDs are assigned automatically.
  - Text arguments containing spaces must be wrapped in quotes.
`)
	return
	}else if len(os.Args)<2 {
		fmt.Println(`No application or command provided.

Run 'help' to see available commands.`)
	return
	}else if os.Args[1]!="notes"{
		fmt.Printf(`Unknown application: %s

Run 'help' to see available commands.`,os.Args[1])
	return
	}



	var application string = os.Args[1]
	var command string = os.Args[2]
	switch command{
	case "add":
		if len(os.Args)<3{
			fmt.Println("Please provide note text and try again.")
		}
		commands.Add(os.Args[2])
	case "list":
		notes := commands.List()
		if len(notes)==0{
			fmt.Println("No notes")
			return
		}
		for _,note:=range notes{
			fmt.Printf("[%d] %s\n",note.ID,note.Text)
		}
	}
}

