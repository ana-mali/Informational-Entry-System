package main

import (
	"fmt"
	"os"
	"NotesApp/cmd/cli/handlers"
)
func main() {
	if len(os.Args) < 2 ||len(os.Args) < 3{
		printHelp("No command provided.")
		return
	}
	if len(os.Args) < 4 {
		printHelp("")
		return
	}
	app := os.Args[1]
	resource := os.Args[2]
	command := os.Args[3]
	args := os.Args[4:]

	if app != "app" {
		printHelp("Unknown application")
		return
	}

	switch resource {
	case "notes":
		handlers.HandleNotes(command, args)
	case "lists":
		handlers.HandleLists(command, args)
	case "tasks":
		handlers.HandleTasks(command, args)
	default:
		fmt.Println("Unknown resource:", resource)
	}
}

func printHelp(msg string) {
	if msg != "" {
		fmt.Println(msg)
		fmt.Println()
	}

	fmt.Println(`Personal Information Organizer

Usage:
	./cmd/cli app <domain> <command> [arguments]
Domains:
	notes		Manage personal notes
	list		Manage lists and list items 
	task		Manage tasks

Commands:
  <domain> add <text>        			Add a new domain
  <domain> list              			List all data from domain 
  <domain> delete <id>       			Delete a domain
  help            		     			Show this help
  list item add <ListID> <text>			Add a new item to a list
  list item delete <ListID> <ItemID>	Delete an item from a list

Examples:
  app notes add "Make a list"
  app notes list 
  app notes delete 2
  app list create "Groceries"
  app list item add 1 "Buy milk"
  app task add "Submit Report" --priority high --due 2026-02-01
  app task delete 1
  app list item remove 1 3

`)
}